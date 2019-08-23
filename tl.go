package mtproto

import (
	"math/big"
)

type TL interface {
	encode() []byte
}

type TlMsgContainer struct {
	Items []TlMtMessage
}

type TlMtMessage struct {
	MsgId int64
	SeqNo int32
	Size  int32
	Data  interface{}
}

type TlReqPq struct {
	Nonce []byte
}

type TlPQInnerData struct {
	pq          *big.Int
	p           *big.Int
	q           *big.Int
	nonce       []byte
	serverNonce []byte
	newNonce    []byte
}
type TlReqDHParams struct {
	nonce       []byte
	serverNonce []byte
	p           *big.Int
	q           *big.Int
	fp          uint64
	encData     []byte
}
type TlClientDHInnerData struct {
	nonce       []byte
	serverNonce []byte
	retry       int64
	gB          *big.Int
}
type TlSetClientDHParams struct {
	nonce       []byte
	serverNonce []byte
	encData     []byte
}
type TlRespq struct {
	nonce        []byte
	serverNonce  []byte
	pq           *big.Int
	fingerprints []int64
}

type TlServerDHParamsOk struct {
	nonce           []byte
	serverNonce     []byte
	encryptedAnswer []byte
}

type TlServerDHInnerData struct {
	nonce       []byte
	serverNonce []byte
	g           int32
	dhPrime     *big.Int
	gA          *big.Int
	serverTime  int32
}

type TlNewSessionCreated struct {
	firstMsgId int64
	uniqueId   int64
	serverSalt []byte
}

type TlBadServerSalt struct {
	badMsgId      int64
	badMsgSeqno   int32
	errorCode     int32
	newServerSalt []byte
}

type TlCrcBadMsgNotification struct {
	badMsgId    int64
	badMsgSeqno int32
	errorCode   int32
}

type TlMsgsAck struct {
	msgIds []int64
}

type TlRpcResult struct {
	reqMsgId int64
	obj      interface{}
}

type TlRpcError struct {
	errorCode    int32
	errorMessage string
}

type TlDHGenOk struct {
	nonce         []byte
	serverNonce   []byte
	newNonceHash1 []byte
}

type TlPing struct {
	pingId int64
}

type TlPong struct {
	msgId  int64
	pingId int64
}
