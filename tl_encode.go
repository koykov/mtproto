package mtproto

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
	"time"
)

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	return ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
}

type EncodeBuf struct {
	buf []byte
}

func NewEncodeBuf(cap int) *EncodeBuf {
	return &EncodeBuf{make([]byte, 0, cap)}
}

func (e *EncodeBuf) Int(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *EncodeBuf) UInt(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *EncodeBuf) Long(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *EncodeBuf) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
}

func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

func (e *EncodeBuf) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.buf = append(e.buf, res...)
}

func (e *EncodeBuf) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *EncodeBuf) VectorInt(v []int32) {
	x := make([]byte, 4+4+len(v)*4)
	binary.LittleEndian.PutUint32(x, CrcVector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint32(x[i:], uint32(v))
		i += 4
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorLong(v []int64) {
	x := make([]byte, 4+4+len(v)*8)
	binary.LittleEndian.PutUint32(x, CrcVector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint64(x[i:], uint64(v))
		i += 8
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorString(v []string) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, CrcVector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.String(v)
	}
}

func (e *EncodeBuf) Vector(v []TL) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, CrcVector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.buf = append(e.buf, v.encode()...)
	}
}

func (e TlMsgContainer) encode() []byte          { return nil }
func (e TlRespq) encode() []byte                 { return nil }
func (e TlServerDHParamsOk) encode() []byte      { return nil }
func (e TlServerDHInnerData) encode() []byte     { return nil }
func (e TlDHGenOk) encode() []byte               { return nil }
func (e TlRpcResult) encode() []byte             { return nil }
func (e TlRpcError) encode() []byte              { return nil }
func (e TlNewSessionCreated) encode() []byte     { return nil }
func (e TlBadServerSalt) encode() []byte         { return nil }
func (e TlCrcBadMsgNotification) encode() []byte { return nil }

func (e TlReqPq) encode() []byte {
	x := NewEncodeBuf(20)
	x.UInt(CrcReqPq)
	x.Bytes(e.Nonce)
	return x.buf
}

func (e TlPQInnerData) encode() []byte {
	x := NewEncodeBuf(256)
	x.UInt(CrcPQInnerData)
	x.BigInt(e.pq)
	x.BigInt(e.p)
	x.BigInt(e.q)
	x.Bytes(e.nonce)
	x.Bytes(e.serverNonce)
	x.Bytes(e.newNonce)
	return x.buf
}

func (e TlReqDHParams) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcReqDHParams)
	x.Bytes(e.nonce)
	x.Bytes(e.serverNonce)
	x.BigInt(e.p)
	x.BigInt(e.q)
	x.Long(int64(e.fp))
	x.StringBytes(e.encData)
	return x.buf
}

func (e TlClientDHInnerData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CrcClientDHInnerData)
	x.Bytes(e.nonce)
	x.Bytes(e.serverNonce)
	x.Long(e.retry)
	x.BigInt(e.gB)
	return x.buf
}

func (e TlSetClientDHParams) encode() []byte {
	x := NewEncodeBuf(256)
	x.UInt(CrcSetClientDHParams)
	x.Bytes(e.nonce)
	x.Bytes(e.serverNonce)
	x.StringBytes(e.encData)
	return x.buf
}

func (e TlPing) encode() []byte {
	x := NewEncodeBuf(32)
	x.UInt(CrcPing)
	x.Long(e.pingId)
	return x.buf
}

func (e TlPong) encode() []byte {
	x := NewEncodeBuf(32)
	x.UInt(CrcPong)
	x.Long(e.msgId)
	x.Long(e.pingId)
	return x.buf
}

func (e TlMsgsAck) encode() []byte {
	x := NewEncodeBuf(64)
	x.UInt(CrcMsgsAck)
	x.VectorLong(e.msgIds)
	return x.buf
}
