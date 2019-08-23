package mtproto

import (
	"errors"
	"fmt"
	"github.com/koykov/telegram/config"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type MTProto struct {
	addr      string
	conn      *net.TCPConn
	f         *os.File
	queueSend chan packetToSend
	stopSend  chan struct{}
	stopRead  chan struct{}
	stopPing  chan struct{}
	allDone   chan struct{}

	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	encrypted   bool
	sessionId   int64

	mutex        *sync.Mutex
	lastSeqNo    int32
	msgsIdToAck  map[int64]packetToSend
	msgsIdToResp map[int64]chan TL
	seqNo        int32
	msgId        int64

	dclist map[int32]string
}

type packetToSend struct {
	msg  TL
	resp chan TL
}

var (
	conf *config.TGConfig
)

func NewMTProto(cnf *config.TGConfig) (*MTProto, error) {
	conf = cnf

	var err error
	m := MTProto{}

	m.f, err = os.OpenFile(conf.SecretKey, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	err = m.readData()
	if err == nil {
		m.encrypted = true
	} else {
		m.addr = conf.ProdAddress
		m.encrypted = false
	}
	rand.Seed(time.Now().UnixNano())
	m.sessionId = rand.Int63()

	return &m, nil
}

func (m *MTProto) Connect() error {
	var err error
	var tcpAddr *net.TCPAddr

	// connect
	tcpAddr, err = net.ResolveTCPAddr("tcp", m.addr)
	if err != nil {
		return err
	}
	m.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	_, err = m.conn.Write([]byte{0xef})
	if err != nil {
		return err
	}

	// get new authKey if need
	if !m.encrypted {
		err = m.makeAuthKey()
		if err != nil {
			return err
		}
	}

	// start goroutines
	m.queueSend = make(chan packetToSend, 64)
	m.stopSend = make(chan struct{}, 1)
	m.stopRead = make(chan struct{}, 1)
	m.stopPing = make(chan struct{}, 1)
	m.allDone = make(chan struct{}, 3)
	m.msgsIdToAck = make(map[int64]packetToSend)
	m.msgsIdToResp = make(map[int64]chan TL)
	m.mutex = &sync.Mutex{}
	go m.sendRoutine()
	go m.readRoutine()

	var resp chan TL
	var x TL

	// (help_getConfig)
	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{
		TLInvokeWithLayer{
			layer,
			TLInitConnection{
				conf.ApiId,
				"Unknown",
				runtime.GOOS + "/" + runtime.GOARCH,
				"0.0.4",
				"en",
				TLHelpGetConfig{},
			},
		},
		resp,
	}
	x = <-resp
	switch x.(type) {
	case TLConfig:
		m.dclist = make(map[int32]string, 5)
		for _, v := range x.(TLConfig).DcOptions {
			v := v.(TLDcOption)
			m.dclist[v.Id] = fmt.Sprintf("%s:%d", v.IpAddress, v.Port)
		}
	default:
		return fmt.Errorf("got: %T", x)
	}

	// start keepalive pinging
	go m.pingRoutine()

	return nil
}

func (m *MTProto) reconnect(newaddr string) error {
	var err error

	// stop ping routine
	m.stopPing <- struct{}{}
	close(m.stopPing)

	// stop send routine
	m.stopSend <- struct{}{}
	close(m.stopSend)

	// stop read routine
	m.stopRead <- struct{}{}
	close(m.stopRead)

	<-m.allDone
	<-m.allDone
	<-m.allDone

	// close send queue
	close(m.queueSend)

	// close connection
	err = m.conn.Close()
	if err != nil {
		return err
	}

	// renew connection
	m.encrypted = false
	m.addr = newaddr
	err = m.Connect()
	return err
}

func (m *MTProto) AuthSendCode(phoneNumber string) (*TLAuthSentCode, error) {
	var authSentCode TLAuthSentCode

	flag := true
	for flag {
		resp := make(chan TL, 1)
		m.queueSend <- packetToSend{
			msg: TLAuthSendCode{
				phoneNumber,
				0,
				conf.ApiId,
				conf.ApiHash,
				"en",
			},
			resp: resp,
		}
		x := <-resp
		switch x.(type) {
		case TLAuthSentCode:
			authSentCode = x.(TLAuthSentCode)
			flag = false
		case TLAuthSentAppCode:
			authSentAppCode := x.(TLAuthSentCode)
			authSentCode.PhoneRegistered = authSentAppCode.PhoneRegistered
			authSentCode.PhoneCodeHash = authSentAppCode.PhoneCodeHash
			authSentCode.SendCallTimeout = authSentAppCode.SendCallTimeout
			authSentCode.IsPassword = authSentAppCode.IsPassword
			flag = false
		case TlRpcError:
			x := x.(TlRpcError)
			if x.errorCode != 303 {
				return nil, fmt.Errorf("RPC error_code: %d", x.errorCode)
			}
			var newDc int32
			n, _ := fmt.Sscanf(x.errorMessage, "PHONE_MIGRATE_%d", &newDc)
			if n != 1 {
				n, _ := fmt.Sscanf(x.errorMessage, "NETWORK_MIGRATE_%d", &newDc)
				if n != 1 {
					return nil, fmt.Errorf("RPC error_string: %s", x.errorMessage)
				}
			}

			newDcAddr, ok := m.dclist[newDc]
			if !ok {
				return nil, fmt.Errorf("wrong DC index: %d", newDc)
			}
			err := m.reconnect(newDcAddr)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("got: %T", x)
		}

	}

	return &authSentCode, nil
}

func (m *MTProto) AuthSignIn(phoneNumber string, authSentCode *TLAuthSentCode, code int) (*TLUserSelf, error) {
	if toBool(authSentCode.PhoneRegistered) {
		resp := make(chan TL, 1)
		m.queueSend <- packetToSend{
			TLAuthSignIn{
				phoneNumber,
				authSentCode.PhoneCodeHash,
				strconv.Itoa(code),
			},
			resp,
		}
		x := <-resp
		auth, ok := x.(TLAuthAuthorization)
		if !ok {
			return nil, fmt.Errorf("RPC: %#v", x)
		}
		userSelf := auth.User.(TLUserSelf)
		return &userSelf, nil
	} else {
		return nil, errors.New("cannot sign in")
	}
}

func (m *MTProto) GetContacts() error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{TLContactsGetContacts{""}, resp}
	x := <-resp
	list, ok := x.(TLContactsContacts)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	contacts := make(map[int32]TLUserContact)
	for _, v := range list.Users {
		if v, ok := v.(TLUserContact); ok {
			contacts[v.Id] = v
		}
	}
	fmt.Printf(
		"\033[33m\033[1m%10s    %10s    %-30s    %-20s\033[0m\n",
		"id", "mutual", "name", "username",
	)
	for _, v := range list.Contacts {
		v := v.(TLContact)
		fmt.Printf(
			"%10d    %10t    %-30s    %-20s\n",
			v.UserId,
			toBool(v.Mutual),
			fmt.Sprintf("%s %s", contacts[v.UserId].FirstName, contacts[v.UserId].LastName),
			contacts[v.UserId].Username,
		)
	}

	return nil
}

func (m *MTProto) SendMessage(userId int32, msg string) error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TLMessagesSendMessage{
			TLInputPeerContact{userId},
			msg,
			rand.Int63(),
		},
		resp,
	}
	x := <-resp
	_, ok := x.(TLMessagesSentMessage)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	return nil
}

func (m *MTProto) pingRoutine() {
	for {
		select {
		case <-m.stopPing:
			m.allDone <- struct{}{}
			return
		case <-time.After(60 * time.Second):
			m.queueSend <- packetToSend{TlPing{0xCADACADA}, nil}
		}
	}
}

func (m *MTProto) sendRoutine() {
	for x := range m.queueSend {
		err := m.sendPacket(x.msg, x.resp)
		if err != nil {
			fmt.Println("SendRoutine:", err)
			os.Exit(2)
		}
	}

	m.allDone <- struct{}{}
}

func (m *MTProto) readRoutine() {
	for {
		data, err := m.read(m.stopRead)
		if err != nil {
			fmt.Println("ReadRoutine:", err)
			os.Exit(2)
		}
		if data == nil {
			m.allDone <- struct{}{}
			return
		}

		m.process(m.msgId, m.seqNo, data)
	}

}

func (m *MTProto) process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case TlMsgContainer:
		data := data.(TlMsgContainer).Items
		for _, v := range data {
			m.process(v.MsgId, v.SeqNo, v.Data)
		}

	case TlBadServerSalt:
		data := data.(TlBadServerSalt)
		m.serverSalt = data.newServerSalt
		_ = m.saveData()
		m.mutex.Lock()
		for k, v := range m.msgsIdToAck {
			delete(m.msgsIdToAck, k)
			m.queueSend <- v
		}
		m.mutex.Unlock()

	case TlNewSessionCreated:
		data := data.(TlNewSessionCreated)
		m.serverSalt = data.serverSalt
		_ = m.saveData()

	case TlPing:
		data := data.(TlPing)
		m.queueSend <- packetToSend{TlPong{msgId, data.pingId}, nil}

	case TlPong:
		// (ignore)

	case TlMsgsAck:
		data := data.(TlMsgsAck)
		m.mutex.Lock()
		for _, v := range data.msgIds {
			delete(m.msgsIdToAck, v)
		}
		m.mutex.Unlock()

	case TlRpcResult:
		data := data.(TlRpcResult)
		x := m.process(msgId, seqNo, data.obj)
		m.mutex.Lock()
		v, ok := m.msgsIdToResp[data.reqMsgId]
		if ok {
			v <- x.(TL)
			close(v)
			delete(m.msgsIdToResp, data.reqMsgId)
		}
		delete(m.msgsIdToAck, data.reqMsgId)
		m.mutex.Unlock()

	default:
		return data

	}

	if (seqNo & 1) == 1 {
		m.queueSend <- packetToSend{TlMsgsAck{[]int64{msgId}}, nil}
	}

	return nil
}

func (m *MTProto) saveData() (err error) {
	m.encrypted = true

	b := NewEncodeBuf(1024)
	b.StringBytes(m.authKey)
	b.StringBytes(m.authKeyHash)
	b.StringBytes(m.serverSalt)
	b.String(m.addr)

	err = m.f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = m.f.WriteAt(b.buf, 0)
	if err != nil {
		return err
	}

	return nil
}

func (m *MTProto) readData() (err error) {
	b := make([]byte, 1024*4)
	n, err := m.f.ReadAt(b, 0)
	if n <= 0 {
		return errors.New("new session")
	}

	d := NewDecodeBuf(b)
	m.authKey = d.StringBytes()
	m.authKeyHash = d.StringBytes()
	m.serverSalt = d.StringBytes()
	m.addr = d.String()

	if d.err != nil {
		return d.err
	}

	return nil
}
