package sessionstate

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAuthKeyStore struct {
	key        *tg.AuthKeyInfo
	futureSalt *tg.FutureSalt
	expiresIn  int32
	setCalls   int
}

func (f *fakeAuthKeyStore) QueryAuthKey(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	return f.key, nil
}

func (f *fakeAuthKeyStore) SetAuthKey(ctx context.Context, authKey *tg.AuthKeyInfo, futureSalt *tg.FutureSalt, expiresIn int32) error {
	f.key = authKey
	f.futureSalt = futureSalt
	f.expiresIn = expiresIn
	f.setCalls++
	return nil
}

func (f *fakeAuthKeyStore) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.FutureSalts, error) {
	return nil, nil
}

func TestAuthHandshakeFullFlow(t *testing.T) {
	store := &fakeAuthKeyStore{}
	manager := NewHandshakeManager(store)
	nonce := testInt128(1)

	resPQMsg := handlePlainForTest(t, manager, 100, encodeTL(t, &mt.TLReqPqMulti{Nonce: nonce}))
	if resPQMsg.MsgId <= 100 {
		t.Fatalf("resPQ msg_id = %d, want > request", resPQMsg.MsgId)
	}
	resPQ := decodeBodyAs[*mt.TLResPQ](t, resPQMsg.Body)
	if resPQ.Nonce != nonce || len(resPQ.ServerPublicKeyFingerprints) == 0 {
		t.Fatalf("resPQ = %#v", resPQ)
	}

	newNonce := testInt256(2)
	reqDH := &mt.TLReqDHParams{
		Nonce:                nonce,
		ServerNonce:          resPQ.ServerNonce,
		P:                    string(handshakeP),
		Q:                    string(handshakeQ),
		PublicKeyFingerprint: resPQ.ServerPublicKeyFingerprints[0],
		EncryptedData:        string(encryptPQInnerForTest(t, manager, nonce, resPQ.ServerNonce, newNonce)),
	}
	serverDHMsg := handlePlainForTest(t, manager, 200, encodeTL(t, reqDH))
	if serverDHMsg.MsgId <= 200 {
		t.Fatalf("server_DH_params_ok msg_id = %d, want > request", serverDHMsg.MsgId)
	}
	serverDH := decodeBodyAs[*mt.TLServerDHParamsOk](t, serverDHMsg.Body)
	serverInner := decryptServerDHForTest(t, newNonce, resPQ.ServerNonce, []byte(serverDH.EncryptedAnswer))
	if serverInner.Nonce != nonce || serverInner.ServerNonce != resPQ.ServerNonce {
		t.Fatalf("server_DH_inner_data = %#v", serverInner)
	}

	clientB := big.NewInt(17)
	gB := new(big.Int).Exp(big.NewInt(int64(serverInner.G)), clientB, new(big.Int).SetBytes([]byte(serverInner.DhPrime))).Bytes()
	setClient := &mt.TLSetClientDHParams{
		Nonce:         nonce,
		ServerNonce:   resPQ.ServerNonce,
		EncryptedData: string(encryptClientDHForTest(t, nonce, resPQ.ServerNonce, newNonce, gB)),
	}
	dhGenMsg := handlePlainForTest(t, manager, 300, encodeTL(t, setClient))
	if dhGenMsg.MsgId <= 300 {
		t.Fatalf("dh_gen_ok msg_id = %d, want > request", dhGenMsg.MsgId)
	}
	dhGen := decodeBodyAs[*mt.TLDhGenOk](t, dhGenMsg.Body)
	if dhGen.Nonce != nonce || dhGen.ServerNonce != resPQ.ServerNonce {
		t.Fatalf("dh_gen_ok = %#v", dhGen)
	}
	if store.setCalls != 1 || store.key == nil || len(store.key.AuthKey) != 256 || store.futureSalt == nil {
		t.Fatalf("saved key calls=%d key=%#v salt=%#v", store.setCalls, store.key, store.futureSalt)
	}
}

func TestAuthHandshakeAcceptsLegacyReqPq(t *testing.T) {
	store := &fakeAuthKeyStore{}
	manager := NewHandshakeManager(store)
	nonce := testInt128(3)

	resPQMsg := handlePlainForTest(t, manager, 100, encodeTL(t, &mt.TLReqPq{Nonce: nonce}))
	resPQ := decodeBodyAs[*mt.TLResPQ](t, resPQMsg.Body)
	if resPQ.Nonce != nonce || len(resPQ.ServerPublicKeyFingerprints) == 0 {
		t.Fatalf("resPQ = %#v", resPQ)
	}
}

func TestAuthHandshakeAcceptsPQInnerDataDc(t *testing.T) {
	store := &fakeAuthKeyStore{}
	manager := NewHandshakeManager(store)
	nonce := testInt128(4)

	resPQMsg := handlePlainForTest(t, manager, 100, encodeTL(t, &mt.TLReqPq{Nonce: nonce}))
	resPQ := decodeBodyAs[*mt.TLResPQ](t, resPQMsg.Body)
	newNonce := testInt256(5)
	reqDH := &mt.TLReqDHParams{
		Nonce:                nonce,
		ServerNonce:          resPQ.ServerNonce,
		P:                    string(handshakeP),
		Q:                    string(handshakeQ),
		PublicKeyFingerprint: resPQ.ServerPublicKeyFingerprints[0],
		EncryptedData: string(encryptPQInnerObjectForTest(t, manager, &mt.TLPQInnerDataDc{
			Pq:          string(handshakePQ),
			P:           string(handshakeP),
			Q:           string(handshakeQ),
			Nonce:       nonce,
			ServerNonce: resPQ.ServerNonce,
			NewNonce:    newNonce,
			Dc:          1,
		})),
	}

	serverDHMsg := handlePlainForTest(t, manager, 200, encodeTL(t, reqDH))
	serverDH := decodeBodyAs[*mt.TLServerDHParamsOk](t, serverDHMsg.Body)
	if serverDH.Nonce != nonce || serverDH.ServerNonce != resPQ.ServerNonce {
		t.Fatalf("server_DH_params_ok = %#v", serverDH)
	}
}

func TestAuthHandshakeRejectsWrongNonce(t *testing.T) {
	store := &fakeAuthKeyStore{}
	manager := NewHandshakeManager(store)
	nonce := testInt128(1)

	resPQMsg := handlePlainForTest(t, manager, 100, encodeTL(t, &mt.TLReqPqMulti{Nonce: nonce}))
	resPQ := decodeBodyAs[*mt.TLResPQ](t, resPQMsg.Body)
	reqDH := &mt.TLReqDHParams{
		Nonce:                testInt128(9),
		ServerNonce:          resPQ.ServerNonce,
		P:                    string(handshakeP),
		Q:                    string(handshakeQ),
		PublicKeyFingerprint: resPQ.ServerPublicKeyFingerprints[0],
		EncryptedData:        string(encryptPQInnerForTest(t, manager, nonce, resPQ.ServerNonce, testInt256(2))),
	}

	if _, err := manager.HandlePlain(context.Background(), gmtproto.PlainMessage{MsgId: 200, Body: encodeTL(t, reqDH)}); err == nil {
		t.Fatal("HandlePlain() error is nil")
	}
	if store.setCalls != 0 {
		t.Fatalf("SetAuthKey calls = %d, want 0", store.setCalls)
	}
}

func TestAuthHandshakeKeepsInterleavedClientState(t *testing.T) {
	store := &fakeAuthKeyStore{}
	manager := NewHandshakeManager(store)
	nonce1 := testInt128(1)
	nonce2 := testInt128(20)

	resPQMsg1 := handlePlainForTest(t, manager, 100, encodeTL(t, &mt.TLReqPqMulti{Nonce: nonce1}))
	resPQ1 := decodeBodyAs[*mt.TLResPQ](t, resPQMsg1.Body)
	_ = handlePlainForTest(t, manager, 101, encodeTL(t, &mt.TLReqPqMulti{Nonce: nonce2}))

	reqDH1 := &mt.TLReqDHParams{
		Nonce:                nonce1,
		ServerNonce:          resPQ1.ServerNonce,
		P:                    string(handshakeP),
		Q:                    string(handshakeQ),
		PublicKeyFingerprint: resPQ1.ServerPublicKeyFingerprints[0],
		EncryptedData:        string(encryptPQInnerForTest(t, manager, nonce1, resPQ1.ServerNonce, testInt256(2))),
	}
	serverDHMsg1 := handlePlainForTest(t, manager, 200, encodeTL(t, reqDH1))
	serverDH1 := decodeBodyAs[*mt.TLServerDHParamsOk](t, serverDHMsg1.Body)
	if serverDH1.Nonce != nonce1 || serverDH1.ServerNonce != resPQ1.ServerNonce {
		t.Fatalf("server_DH_params_ok = %#v", serverDH1)
	}
}

func TestAuthHandshakeExpiresStaleState(t *testing.T) {
	store := &fakeAuthKeyStore{}
	manager := NewHandshakeManager(store)
	nonce := testInt128(1)

	resPQMsg := handlePlainForTest(t, manager, 100, encodeTL(t, &mt.TLReqPqMulti{Nonce: nonce}))
	resPQ := decodeBodyAs[*mt.TLResPQ](t, resPQMsg.Body)
	manager.mu.Lock()
	state := manager.states[resPQ.ServerNonce]
	state.expiresAt = time.Now().Add(-time.Second)
	manager.states[resPQ.ServerNonce] = state
	manager.mu.Unlock()

	reqDH := &mt.TLReqDHParams{
		Nonce:                nonce,
		ServerNonce:          resPQ.ServerNonce,
		P:                    string(handshakeP),
		Q:                    string(handshakeQ),
		PublicKeyFingerprint: resPQ.ServerPublicKeyFingerprints[0],
		EncryptedData:        string(encryptPQInnerForTest(t, manager, nonce, resPQ.ServerNonce, testInt256(2))),
	}
	if _, err := manager.HandlePlain(context.Background(), gmtproto.PlainMessage{MsgId: 200, Body: encodeTL(t, reqDH)}); err == nil {
		t.Fatal("HandlePlain() error is nil")
	}

	manager.mu.Lock()
	_, ok := manager.states[resPQ.ServerNonce]
	manager.mu.Unlock()
	if ok {
		t.Fatal("expired handshake state was not pruned")
	}
}

func handlePlainForTest(t *testing.T, manager *HandshakeManager, msgID int64, body []byte) gmtproto.PlainMessage {
	t.Helper()
	resp, err := manager.HandlePlain(context.Background(), gmtproto.PlainMessage{MsgId: msgID, Body: body})
	if err != nil {
		t.Fatalf("HandlePlain() error = %v", err)
	}
	msg, err := gmtproto.DecodePlainMessage(resp)
	if err != nil {
		t.Fatalf("DecodePlainMessage() error = %v", err)
	}
	return msg
}

func encryptPQInnerForTest(t *testing.T, manager *HandshakeManager, nonce, serverNonce bin.Int128, newNonce bin.Int256) []byte {
	t.Helper()
	return encryptPQInnerObjectForTest(t, manager, &mt.TLPQInnerData{
		Pq:          string(handshakePQ),
		P:           string(handshakeP),
		Q:           string(handshakeQ),
		Nonce:       nonce,
		ServerNonce: serverNonce,
		NewNonce:    newNonce,
	})
}

func encryptPQInnerObjectForTest(t *testing.T, manager *HandshakeManager, obj iface.TLObject) []byte {
	t.Helper()
	inner := encodeTL(t, obj)
	data := make([]byte, 192)
	copy(data, inner)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	modulus := rsaPrivateKeyForTest(t).N
	for seed := byte(1); seed != 0; seed++ {
		tempKey := bytes.Repeat([]byte{seed}, 32)
		hash := crypto.Sha256Digest(append(append([]byte(nil), tempKey...), data...))
		withHash := append(data, hash...)
		encrypted, err := crypto.NewAES256IGECryptor(tempKey, zeroIV).Encrypt(withHash)
		if err != nil {
			t.Fatalf("encrypt pq inner: %v", err)
		}
		tempKeyXor := append([]byte(nil), tempKey...)
		encryptedHash := crypto.Sha256Digest(encrypted)
		for i := range tempKeyXor {
			tempKeyXor[i] ^= encryptedHash[i]
		}
		block := append(tempKeyXor, encrypted...)
		if new(big.Int).SetBytes(block).Cmp(modulus) < 0 {
			return rsaEncryptForTest(t, block)
		}
	}
	t.Fatal("failed to build RSA payload smaller than modulus")
	return nil
}

func decryptServerDHForTest(t *testing.T, newNonce bin.Int256, serverNonce bin.Int128, encrypted []byte) *mt.TLServerDHInnerData {
	t.Helper()
	key, iv := tmpAESKeyIV(newNonce, serverNonce)
	decrypted, err := crypto.NewAES256IGECryptor(key, iv).Decrypt(encrypted)
	if err != nil {
		t.Fatalf("decrypt server dh: %v", err)
	}
	if !checkSha1Prefix(decrypted) {
		t.Fatalf("server dh sha1 prefix mismatch")
	}
	return decodeRawAs[*mt.TLServerDHInnerData](t, decrypted[20:])
}

func encryptClientDHForTest(t *testing.T, nonce, serverNonce bin.Int128, newNonce bin.Int256, gB []byte) []byte {
	t.Helper()
	inner := encodeTL(t, &mt.TLClientDHInnerData{
		Nonce:       nonce,
		ServerNonce: serverNonce,
		RetryId:     0,
		GB:          string(gB),
	})
	data := make([]byte, 20+len(inner))
	sum := sha1.Sum(inner)
	copy(data, sum[:])
	copy(data[20:], inner)
	for len(data)%16 != 0 {
		data = append(data, 0)
	}
	key, iv := tmpAESKeyIV(newNonce, serverNonce)
	encrypted, err := crypto.NewAES256IGECryptor(key, iv).Encrypt(data)
	if err != nil {
		t.Fatalf("encrypt client dh: %v", err)
	}
	return encrypted
}

func testInt128(seed byte) bin.Int128 {
	var v bin.Int128
	for i := range v {
		v[i] = seed + byte(i)
	}
	return v
}

func testInt256(seed byte) bin.Int256 {
	var v bin.Int256
	for i := range v {
		v[i] = seed + byte(i)
	}
	return v
}

func decodeBodyAs[T any](t *testing.T, body []byte) T {
	t.Helper()
	obj, err := iface.DecodeObject(bin.NewDecoder(body))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	got, ok := obj.(T)
	if !ok {
		t.Fatalf("DecodeObject() = %T", obj)
	}
	return got
}

func decodeRawAs[T any](t *testing.T, body []byte) T {
	t.Helper()
	return decodeBodyAs[T](t, body)
}

func encodeTL(t *testing.T, obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	t.Helper()
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 0); err != nil {
		x.Reset()
		if err2 := obj.Encode(x, 224); err2 != nil {
			t.Fatalf("Encode(%T) error = %v", obj, err)
		}
	}
	return append([]byte(nil), x.Bytes()...)
}

func rsaEncryptForTest(t *testing.T, payload []byte) []byte {
	t.Helper()
	key := rsaPrivateKeyForTest(t)
	ciphertext := new(big.Int).Exp(new(big.Int).SetBytes(payload), big.NewInt(int64(key.E)), key.N).Bytes()
	out := make([]byte, 256)
	copy(out[256-len(ciphertext):], ciphertext)
	return out
}

func rsaPrivateKeyForTest(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	block, _ := pem.Decode(defaultRSAPrivateKey)
	if block == nil {
		t.Fatal("default RSA private key PEM decode failed")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("ParsePKCS1PrivateKey() error = %v", err)
	}
	return key
}
