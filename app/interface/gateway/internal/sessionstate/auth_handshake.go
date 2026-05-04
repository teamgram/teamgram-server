package sessionstate

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"
	"time"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var (
	handshakePQ = []byte{0x17, 0xed, 0x48, 0x94, 0x1a, 0x08, 0xf9, 0x81}
	handshakeP  = []byte{0x49, 0x4c, 0x55, 0x3b}
	handshakeQ  = []byte{0x53, 0x91, 0x10, 0x73}
	handshakeG  = int32(3)
	handshakeP2 = []byte{
		0xc7, 0x1c, 0xae, 0xb9, 0xc6, 0xb1, 0xc9, 0x04, 0x8e, 0x6c, 0x52, 0x2f,
		0x70, 0xf1, 0x3f, 0x73, 0x98, 0x0d, 0x40, 0x23, 0x8e, 0x3e, 0x21, 0xc1,
		0x49, 0x34, 0xd0, 0x37, 0x56, 0x3d, 0x93, 0x0f, 0x48, 0x19, 0x8a, 0x0a,
		0xa7, 0xc1, 0x40, 0x58, 0x22, 0x94, 0x93, 0xd2, 0x25, 0x30, 0xf4, 0xdb,
		0xfa, 0x33, 0x6f, 0x6e, 0x0a, 0xc9, 0x25, 0x13, 0x95, 0x43, 0xae, 0xd4,
		0x4c, 0xce, 0x7c, 0x37, 0x20, 0xfd, 0x51, 0xf6, 0x94, 0x58, 0x70, 0x5a,
		0xc6, 0x8c, 0xd4, 0xfe, 0x6b, 0x6b, 0x13, 0xab, 0xdc, 0x97, 0x46, 0x51,
		0x29, 0x69, 0x32, 0x84, 0x54, 0xf1, 0x8f, 0xaf, 0x8c, 0x59, 0x5f, 0x64,
		0x24, 0x77, 0xfe, 0x96, 0xbb, 0x2a, 0x94, 0x1d, 0x5b, 0xcd, 0x1d, 0x4a,
		0xc8, 0xcc, 0x49, 0x88, 0x07, 0x08, 0xfa, 0x9b, 0x37, 0x8e, 0x3c, 0x4f,
		0x3a, 0x90, 0x60, 0xbe, 0xe6, 0x7c, 0xf9, 0xa4, 0xa4, 0xa6, 0x95, 0x81,
		0x10, 0x51, 0x90, 0x7e, 0x16, 0x27, 0x53, 0xb5, 0x6b, 0x0f, 0x6b, 0x41,
		0x0d, 0xba, 0x74, 0xd8, 0xa8, 0x4b, 0x2a, 0x14, 0xb3, 0x14, 0x4e, 0x0e,
		0xf1, 0x28, 0x47, 0x54, 0xfd, 0x17, 0xed, 0x95, 0x0d, 0x59, 0x65, 0xb4,
		0xb9, 0xdd, 0x46, 0x58, 0x2d, 0xb1, 0x17, 0x8d, 0x16, 0x9c, 0x6b, 0xc4,
		0x65, 0xb0, 0xd6, 0xff, 0x9c, 0xa3, 0x92, 0x8f, 0xef, 0x5b, 0x9a, 0xe4,
		0xe4, 0x18, 0xfc, 0x15, 0xe8, 0x3e, 0xbe, 0xa0, 0xf8, 0x7f, 0xa9, 0xff,
		0x5e, 0xed, 0x70, 0x05, 0x0d, 0xed, 0x28, 0x49, 0xf4, 0x7b, 0xf9, 0x59,
		0xd9, 0x56, 0x85, 0x0c, 0xe9, 0x29, 0x85, 0x1f, 0x0d, 0x81, 0x15, 0xf6,
		0x35, 0xb1, 0x05, 0xee, 0x2e, 0x4e, 0x15, 0xd0, 0x4b, 0x24, 0x54, 0xbf,
		0x6f, 0x4f, 0xad, 0xf0, 0x34, 0xb1, 0x04, 0x03, 0x11, 0x9c, 0xd8, 0xe3,
		0xb9, 0x2f, 0xcc, 0x5b,
	}
	zeroIV = make([]byte, 32)
)

const handshakeStateTTL = 2 * time.Minute

type HandshakeManager struct {
	store       repository.AuthKeyStore
	rsa         *crypto.RSACryptor
	fingerprint int64
	mu          sync.Mutex
	states      map[bin.Int128]handshakeState
}

type handshakeState struct {
	nonce       bin.Int128
	serverNonce bin.Int128
	newNonce    bin.Int256
	a           []byte
	gA          []byte
	authKeyType int32
	expiresIn   int32
	expiresAt   time.Time
}

type pqInnerData struct {
	pq          string
	p           string
	q           string
	nonce       bin.Int128
	serverNonce bin.Int128
	newNonce    bin.Int256
	authKeyType int32
	expiresIn   int32
}

func NewHandshakeManager(store repository.AuthKeyStore) *HandshakeManager {
	rsa, err := crypto.NewRSACryptorByKeyData(defaultRSAPrivateKey)
	if err != nil {
		panic(err)
	}
	return &HandshakeManager{
		store:       store,
		rsa:         rsa,
		fingerprint: int64(-6205835210776354611),
		states:      make(map[bin.Int128]handshakeState),
	}
}

func (m *HandshakeManager) HandlePlain(ctx context.Context, msg gmtproto.PlainMessage) ([]byte, error) {
	obj, err := iface.DecodeObject(bin.NewDecoder(msg.Body))
	if err != nil {
		return nil, fmt.Errorf("auth handshake decode request: %w", err)
	}

	var response iface.TLObject
	switch req := obj.(type) {
	case *mt.TLReqPq:
		response, err = m.handleReqPQNonce(req.Nonce, "req_pq")
	case *mt.TLReqPqMulti:
		response, err = m.handleReqPQNonce(req.Nonce, "req_pq_multi")
	case *mt.TLReqDHParams:
		response, err = m.handleReqDHParams(req)
	case *mt.TLSetClientDHParams:
		response, err = m.handleSetClientDHParams(ctx, req)
	default:
		err = fmt.Errorf("auth handshake unsupported plain request %T", obj)
	}
	if err != nil {
		return nil, err
	}

	x := bin.NewEncoder()
	defer x.End()
	if err := response.Encode(x, 0); err != nil {
		return nil, fmt.Errorf("auth handshake encode response: %w", err)
	}
	return gmtproto.EncodePlainMessage(gmtproto.PlainMessage{
		MsgId: gmtproto.NextServerMsgId(msg.MsgId),
		Body:  append([]byte(nil), x.Bytes()...),
	})
}

func (m *HandshakeManager) handleReqPQNonce(nonce bin.Int128, op string) (iface.TLObject, error) {
	if nonce.Zero() {
		return nil, fmt.Errorf("auth handshake %s: nonce is zero", op)
	}
	var serverNonce bin.Int128
	copy(serverNonce[:], crypto.GenerateNonce(16))
	now := time.Now()
	m.mu.Lock()
	m.pruneExpiredStatesLocked(now)
	m.states[serverNonce] = handshakeState{nonce: nonce, serverNonce: serverNonce, expiresAt: now.Add(handshakeStateTTL)}
	m.mu.Unlock()
	return mt.MakeTLResPQ(&mt.TLResPQ{
		Nonce:                       nonce,
		ServerNonce:                 serverNonce,
		Pq:                          string(handshakePQ),
		ServerPublicKeyFingerprints: []int64{m.fingerprint},
	}), nil
}

func (m *HandshakeManager) handleReqDHParams(req *mt.TLReqDHParams) (iface.TLObject, error) {
	state, ok := m.handshakeState(req.ServerNonce)
	if !ok {
		return nil, fmt.Errorf("auth handshake req_DH_params: server_nonce mismatch")
	}
	if req.Nonce != state.nonce {
		return nil, fmt.Errorf("auth handshake req_DH_params: nonce mismatch")
	}
	if req.ServerNonce != state.serverNonce {
		return nil, fmt.Errorf("auth handshake req_DH_params: server_nonce mismatch")
	}
	if !bytes.Equal([]byte(req.P), handshakeP) || !bytes.Equal([]byte(req.Q), handshakeQ) {
		return nil, fmt.Errorf("auth handshake req_DH_params: invalid p or q")
	}
	if req.PublicKeyFingerprint != m.fingerprint {
		return nil, fmt.Errorf("auth handshake req_DH_params: invalid fingerprint")
	}

	pqInner, err := m.decryptPQInnerData([]byte(req.EncryptedData))
	if err != nil {
		return nil, err
	}
	if pqInner.nonce != req.Nonce || pqInner.serverNonce != req.ServerNonce {
		return nil, fmt.Errorf("auth handshake req_DH_params: inner nonce mismatch")
	}
	if !bytes.Equal([]byte(pqInner.pq), handshakePQ) || !bytes.Equal([]byte(pqInner.p), handshakeP) || !bytes.Equal([]byte(pqInner.q), handshakeQ) {
		return nil, fmt.Errorf("auth handshake req_DH_params: invalid inner pq")
	}

	a := crypto.GenerateNonce(256)
	gA := new(big.Int).Exp(big.NewInt(int64(handshakeG)), new(big.Int).SetBytes(a), new(big.Int).SetBytes(handshakeP2)).Bytes()
	state.newNonce = pqInner.newNonce
	state.a = a
	state.gA = gA
	state.authKeyType = pqInner.authKeyType
	state.expiresIn = pqInner.expiresIn
	m.setHandshakeState(state)

	inner := mt.MakeTLServerDHInnerData(&mt.TLServerDHInnerData{
		Nonce:       req.Nonce,
		ServerNonce: req.ServerNonce,
		G:           handshakeG,
		DhPrime:     string(handshakeP2),
		GA:          string(gA),
		ServerTime:  int32(time.Now().Unix()),
	})
	encrypted, err := encryptWithTempKey(state.newNonce, state.serverNonce, encodeObject(inner))
	if err != nil {
		return nil, fmt.Errorf("auth handshake req_DH_params: encrypt server dh: %w", err)
	}
	return mt.MakeTLServerDHParamsOk(&mt.TLServerDHParamsOk{
		Nonce:           req.Nonce,
		ServerNonce:     req.ServerNonce,
		EncryptedAnswer: string(encrypted),
	}), nil
}

func (m *HandshakeManager) handleSetClientDHParams(ctx context.Context, req *mt.TLSetClientDHParams) (iface.TLObject, error) {
	state, ok := m.handshakeState(req.ServerNonce)
	if !ok {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: server_nonce mismatch")
	}
	if req.Nonce != state.nonce {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: nonce mismatch")
	}
	if req.ServerNonce != state.serverNonce {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: server_nonce mismatch")
	}

	decrypted, err := decryptWithTempKey(state.newNonce, state.serverNonce, []byte(req.EncryptedData))
	if err != nil {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: decrypt client dh: %w", err)
	}
	if !checkSha1Prefix(decrypted) {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: sha1 mismatch")
	}
	clientDH, ok := mustDecodeObject(decrypted[20:]).(*mt.TLClientDHInnerData)
	if !ok {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: invalid client dh inner data")
	}
	if clientDH.Nonce != req.Nonce || clientDH.ServerNonce != req.ServerNonce {
		return nil, fmt.Errorf("auth handshake set_client_DH_params: inner nonce mismatch")
	}

	authKeyNum := new(big.Int).Exp(new(big.Int).SetBytes([]byte(clientDH.GB)), new(big.Int).SetBytes(state.a), new(big.Int).SetBytes(handshakeP2))
	authKey := make([]byte, 256)
	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())
	authKeyID := calcAuthKeyID(authKey)
	now := int32(time.Now().Unix())
	futureSalt := tg.MakeTLFutureSalt(&tg.FutureSalt{
		ValidSince: now,
		ValidUntil: now + 30*60,
		Salt:       calcServerSalt(state.newNonce, state.serverNonce),
	})
	if m.store != nil {
		if err := m.store.SetAuthKey(ctx, tg.NewAuthKeyInfo(authKeyID, authKey, int(state.authKeyType)), futureSalt, state.expiresIn); err != nil {
			return nil, fmt.Errorf("auth handshake set_client_DH_params: save auth key: %w", err)
		}
	}
	m.deleteHandshakeState(req.ServerNonce)
	return mt.MakeTLDhGenOk(&mt.TLDhGenOk{
		Nonce:         req.Nonce,
		ServerNonce:   req.ServerNonce,
		NewNonceHash1: calcNewNonceHash(state.newNonce, authKey, 1),
	}), nil
}

func (m *HandshakeManager) handshakeState(serverNonce bin.Int128) (handshakeState, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	state, ok := m.states[serverNonce]
	if !ok {
		return handshakeState{}, false
	}
	if time.Now().After(state.expiresAt) {
		delete(m.states, serverNonce)
		return handshakeState{}, false
	}
	return state, ok
}

func (m *HandshakeManager) setHandshakeState(state handshakeState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if state.expiresAt.IsZero() {
		state.expiresAt = time.Now().Add(handshakeStateTTL)
	}
	m.states[state.serverNonce] = state
}

func (m *HandshakeManager) deleteHandshakeState(serverNonce bin.Int128) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.states, serverNonce)
}

func (m *HandshakeManager) pruneExpiredStatesLocked(now time.Time) {
	for serverNonce, state := range m.states {
		if now.After(state.expiresAt) {
			delete(m.states, serverNonce)
		}
	}
}

func (m *HandshakeManager) decryptPQInnerData(encrypted []byte) (*pqInnerData, error) {
	innerData := m.rsa.Decrypt(encrypted)
	if len(innerData) < 256 {
		padded := make([]byte, 256)
		copy(padded[256-len(innerData):], innerData)
		innerData = padded
	}
	if len(innerData) != 256 {
		return nil, fmt.Errorf("auth handshake req_DH_params: invalid encrypted data")
	}
	tempKey := append([]byte(nil), innerData[:32]...)
	hash := crypto.Sha256Digest(innerData[32:])
	for i := range tempKey {
		tempKey[i] ^= hash[i]
	}
	decrypted, err := crypto.NewAES256IGECryptor(tempKey, zeroIV).Decrypt(innerData[32:])
	if err != nil {
		return nil, fmt.Errorf("auth handshake req_DH_params: decrypt pq inner: %w", err)
	}
	if len(decrypted) < 192 {
		return nil, fmt.Errorf("auth handshake req_DH_params: pq inner too short")
	}
	data := append([]byte(nil), decrypted[:192]...)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return normalizePQInnerData(mustDecodeObject(data))
}

func normalizePQInnerData(obj iface.TLObject) (*pqInnerData, error) {
	switch inner := obj.(type) {
	case *mt.TLPQInnerData:
		return &pqInnerData{
			pq:          inner.Pq,
			p:           inner.P,
			q:           inner.Q,
			nonce:       inner.Nonce,
			serverNonce: inner.ServerNonce,
			newNonce:    inner.NewNonce,
			authKeyType: tg.AuthKeyTypePerm,
		}, nil
	case *mt.TLPQInnerDataDc:
		return &pqInnerData{
			pq:          inner.Pq,
			p:           inner.P,
			q:           inner.Q,
			nonce:       inner.Nonce,
			serverNonce: inner.ServerNonce,
			newNonce:    inner.NewNonce,
			authKeyType: tg.AuthKeyTypePerm,
		}, nil
	case *mt.TLPQInnerDataTemp:
		return &pqInnerData{
			pq:          inner.Pq,
			p:           inner.P,
			q:           inner.Q,
			nonce:       inner.Nonce,
			serverNonce: inner.ServerNonce,
			newNonce:    inner.NewNonce,
			authKeyType: tg.AuthKeyTypeTemp,
			expiresIn:   inner.ExpiresIn,
		}, nil
	case *mt.TLPQInnerDataTempDc:
		authKeyType := int32(tg.AuthKeyTypeTemp)
		if inner.Dc < 0 {
			authKeyType = tg.AuthKeyTypeMediaTemp
		}
		return &pqInnerData{
			pq:          inner.Pq,
			p:           inner.P,
			q:           inner.Q,
			nonce:       inner.Nonce,
			serverNonce: inner.ServerNonce,
			newNonce:    inner.NewNonce,
			authKeyType: authKeyType,
			expiresIn:   inner.ExpiresIn,
		}, nil
	default:
		return nil, fmt.Errorf("auth handshake req_DH_params: invalid pq inner object %T", obj)
	}
}

func encodeObject(obj iface.TLObject) []byte {
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 0); err != nil {
		panic(err)
	}
	return append([]byte(nil), x.Bytes()...)
}

func mustDecodeObject(raw []byte) iface.TLObject {
	obj, err := iface.DecodeObject(bin.NewDecoder(raw))
	if err != nil {
		panic(err)
	}
	return obj
}

func encryptWithTempKey(newNonce bin.Int256, serverNonce bin.Int128, body []byte) ([]byte, error) {
	data := make([]byte, 20+len(body))
	sum := sha1.Sum(body)
	copy(data, sum[:])
	copy(data[20:], body)
	for len(data)%16 != 0 {
		data = append(data, 0)
	}
	key, iv := tmpAESKeyIV(newNonce, serverNonce)
	return crypto.NewAES256IGECryptor(key, iv).Encrypt(data)
}

func decryptWithTempKey(newNonce bin.Int256, serverNonce bin.Int128, encrypted []byte) ([]byte, error) {
	key, iv := tmpAESKeyIV(newNonce, serverNonce)
	return crypto.NewAES256IGECryptor(key, iv).Decrypt(encrypted)
}

func tmpAESKeyIV(newNonce bin.Int256, serverNonce bin.Int128) ([]byte, []byte) {
	sha1A := sha1.Sum(append(append([]byte(nil), newNonce[:]...), serverNonce[:]...))
	sha1B := sha1.Sum(append(append([]byte(nil), serverNonce[:]...), newNonce[:]...))
	sha1C := sha1.Sum(append(append([]byte(nil), newNonce[:]...), newNonce[:]...))
	out := make([]byte, 64)
	copy(out, sha1A[:])
	copy(out[20:], sha1B[:])
	copy(out[40:], sha1C[:])
	copy(out[60:], newNonce[:4])
	return out[:32], out[32:]
}

func checkSha1Prefix(data []byte) bool {
	for padding := 0; padding <= 16 && 20+padding <= len(data); padding++ {
		sum := sha1.Sum(data[20 : len(data)-padding])
		if bytes.Equal(sum[:], data[:20]) {
			return true
		}
	}
	return false
}

func calcAuthKeyID(authKey []byte) int64 {
	sum := crypto.Sha1Digest(authKey)
	return int64(binary.LittleEndian.Uint64(sum[12:]))
}

func calcServerSalt(newNonce bin.Int256, serverNonce bin.Int128) int64 {
	var salt int64
	for i := 7; i >= 0; i-- {
		salt <<= 8
		salt |= int64(newNonce[i] ^ serverNonce[i])
	}
	return salt
}

func calcNewNonceHash(newNonce bin.Int256, authKey []byte, b byte) bin.Int128 {
	buf := make([]byte, 0, 73)
	buf = append(buf, newNonce[:]...)
	buf = append(buf, b)
	sha1D := sha1.Sum(authKey)
	buf = append(buf, sha1D[:]...)
	sha1E := sha1.Sum(buf[:41])
	buf = append(buf, sha1E[:]...)
	var out bin.Int128
	copy(out[:], buf[len(buf)-16:])
	return out
}

var defaultRSAPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAvKLEOWTzt9Hn3/9Kdp/RdHcEhzmd8xXeLSpHIIzaXTLJDw8B
hJy1jR/iqeG8Je5yrtVabqMSkA6ltIpgylH///FojMsX1BHu4EPYOXQgB0qOi6kr
08iXZIH9/iOPQOWDsL+Lt8gDG0xBy+sPe/2ZHdzKMjX6O9B4sOsxjFrk5qDoWDri
oJorAJ7eFAfPpOBf2w73ohXudSrJE0lbQ8pCWNpMY8cB9i8r+WBitcvouLDAvmtn
TX7akhoDzmKgpJBYliAY4qA73v7u5UIepE8QgV0jCOhxJCPubP8dg+/PlLLVKyxU
5CdiQtZj2EMy4s9xlNKzX8XezE0MHEa6bQpnFwIDAQABAoIBACd+SGjfyursZoiO
MW/ejAK/PFJ3bKtNI8P++v9Enh8vF8swUBgMmzIdv93jZfnnD1mtT46kU6mXd3fy
FMunGVrjlwkLKET9MC8B5U46Es6T/H4fAA8KCzA+ywefOEnVA5pIsB7dIFFhyNDB
uO8zrBAFfsu+Y1KMlggsZaZGDXB/WVyUJDbEOMZstVx4uNhpcEgKYp28YQMP/yvv
dp4UgnTxXXXpDghzO5iqi5tUWY0p1lH2ii2OZBxEdqdDl7TirorhUDYIivyoe3B5
H30RNBRok/6w7W0WPyY2lSIcjd3cLPte6vx0QfBXVo2A6N9LTKAtAw3iWBp0x9NZ
N5p8OeECgYEA8QywXlM8nH5M7Sg2sMUYBOHA22O26ZPio7rJzcb8dlkV5gVHm+Kl
aDP61Uy8KoYABQ5kFdem/IQAUPepLxmJmiqfbwOIjfajOD3uVAQunFnDCHBWm4Uk
onbpdA5NlT/OUoSjIBemiBR/4CpDK1cEby/sg+EvQaGxqtedEe4xFmcCgYEAyFXe
MyAAOLpzmnCs9NYTTvMPofW8y+kLDodfbskl7M8q6l20VMo/E+g1gQ+65Aah901Z
/LKGi6HpzmHi5q9O2OJtqyI6FVwjXa07M5ueDbHcVKJw4hC9W0oHpMg8hqumPAWF
+MoN/Toy77p5LzoR30WUdhPvOAJPEL1p2a6r29ECgYEAiXfCEVkI5PqGZm2bmv4b
75TLhpJ8WwMSqms48Vi828V8Xpy+NOFxkVargv9rBBk9Y6TMYUSGH9Yr1AEZhBnd
RoVuPUJXmxaACPAQvetQpavvNR3T1od82AZWpvANQMONp7Oqz/+M4mhGcRHJEqti
hQJgsOk4KQbMqvChy/r6FZsCgYEAwyaqgkD9FkXC0UJLqWFUg8bQhqPcGwLUC34h
n8kAUbPpiU5omWQ+mATPAf8xvmkbo81NCJVb7W93U90U7ET/2NSRonCABkiwBtP2
ZKqGB68oA6YNspo960ytL38DPui80aFLxXQGtpPYBKEw5al6uXWNTozSrkvJe3QY
Rb4amdECgYBpGk7zPcK1TbJ++W5fkiory4qOdf0L1Zf0NbML4fY6dIww+dwMVUpq
FbsgCLqimqOFaaECU+LQEFUHHM7zrk7NBf7GzBvQ+qJx8zhJ66sFVox+IirBUyR9
Vh0+z5tIbFbKmYkO06NbeMlq87JexSlocPZtA3HMhEga5/0fHNHsNw==
-----END RSA PRIVATE KEY-----`)
