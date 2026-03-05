// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package gnet

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	httpcodec "github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet/http"
	"github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel"
)

var validPaths = map[string]bool{
	"/api":           true,
	"/apiw1":         true,
	"/apiw1_test":    true,
	"/apiw1_premium": true,
}

func (s *Server) onHttpData(ctx *connContext, c gnet.Conn) gnet.Action {
	if action := ctx.httpCodec.ReadBufferBytes(c); action != gnet.None {
		return action
	}

	body, done, err := ctx.httpCodec.ParseRequest()
	if err != nil {
		logx.Errorf("conn(%s) HTTP parse error: %v", c, err)
		_, _ = c.Write(httpcodec.FormatErrorResponse(400, "Bad Request"))
		return gnet.Close
	}
	if !done {
		return gnet.None
	}

	method := ctx.httpCodec.Method()
	path := ctx.httpCodec.Path()

	// Handle OPTIONS preflight
	if method == "OPTIONS" {
		_, _ = c.Write(httpcodec.FormatCORSResponse())
		ctx.httpCodec.Reset()
		return gnet.None
	}

	// Validate method
	if method != "POST" {
		_, _ = c.Write(httpcodec.FormatErrorResponse(405, "Method Not Allowed"))
		return gnet.Close
	}

	// Validate path
	if !validPaths[path] {
		_, _ = c.Write(httpcodec.FormatErrorResponse(404, "Not Found"))
		return gnet.Close
	}

	// Empty body
	if len(body) < 8 {
		_, _ = c.Write(httpcodec.FormatErrorResponse(400, "Bad Request"))
		return gnet.Close
	}

	authKeyId := int64(binary.LittleEndian.Uint64(body))

	if authKeyId == 0 {
		s.onHttpHandshake(ctx, c, body)
	} else {
		s.onHttpEncryptedMessage(ctx, c, authKeyId, body)
	}

	return gnet.None
}

func (s *Server) onHttpHandshake(ctx *connContext, c gnet.Conn, body []byte) {
	_, obj, err := parseFromIncomingMessage(body[8:])
	if err != nil {
		logx.Errorf("conn(%s) HTTP handshake parse error: %v", c, err)
		_, _ = c.Write(httpcodec.FormatResponse(nil))
		ctx.httpCodec.Reset()
		return
	}

	switch request := obj.(type) {
	case *mtproto.TLReqPq:
		resPQ, err := s.onReqPq(c, request)
		if err != nil {
			logx.Errorf("conn(%s) HTTP onReqPq error: %v", c, err)
			_, _ = c.Write(httpcodec.FormatResponse(nil))
			ctx.httpCodec.Reset()
			return
		}
		s.PutHttpHandshakeState(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		payload := func() []byte {
			x := mtproto.GetEncodeBuf()
			defer mtproto.PutEncodeBuf(x)
			serializeToBuffer(x, mtproto.GenerateMessageId(), resPQ)
			return append([]byte(nil), x.GetBuf()...)
		}()
		_, _ = c.Write(httpcodec.FormatResponse(payload))
		ctx.httpCodec.Reset()

	case *mtproto.TLReqPqMulti:
		resPQ, err := s.onReqPqMulti(c, request)
		if err != nil {
			logx.Errorf("conn(%s) HTTP onReqPqMulti error: %v", c, err)
			_, _ = c.Write(httpcodec.FormatResponse(nil))
			ctx.httpCodec.Reset()
			return
		}
		s.PutHttpHandshakeState(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		payload := func() []byte {
			x := mtproto.GetEncodeBuf()
			defer mtproto.PutEncodeBuf(x)
			serializeToBuffer(x, mtproto.GenerateMessageId(), resPQ)
			return append([]byte(nil), x.GetBuf()...)
		}()
		_, _ = c.Write(httpcodec.FormatResponse(payload))
		ctx.httpCodec.Reset()

	case *mtproto.TLReq_DHParams:
		state := s.GetHttpHandshakeState(request.Nonce)
		if state == nil {
			logx.Errorf("conn(%s) HTTP req_DH_params: state not found", c)
			_, _ = c.Write(httpcodec.FormatResponse(nil))
			ctx.httpCodec.Reset()
			return
		}
		connId := c.ConnId()
		if err := s.pool.Submit(func() {
			serverDHParams, err := s.httpOnReqDHParams(state, request)
			s.eng.Trigger(connId, func(c gnet.Conn) {
				if err != nil {
					logx.Errorf("conn(%s) HTTP onReqDHParams error: %v", c, err)
					_, _ = c.Write(httpcodec.FormatResponse(nil))
				} else {
					reply := func() []byte {
						xr := mtproto.GetEncodeBuf()
						defer mtproto.PutEncodeBuf(xr)
						serializeToBuffer(xr, mtproto.GenerateMessageId(), serverDHParams)
						return append([]byte(nil), xr.GetBuf()...)
					}()
					_, _ = c.Write(httpcodec.FormatResponse(reply))
				}
				if ctx2, _ := c.Context().(*connContext); ctx2 != nil {
					ctx2.httpCodec.Reset()
				}
			})
		}); err != nil {
			logx.Errorf("conn(%s) HTTP pool.Submit error: %v", c, err)
			_, _ = c.Write(httpcodec.FormatResponse(nil))
			ctx.httpCodec.Reset()
		}

	case *mtproto.TLSetClient_DHParams:
		state := s.GetHttpHandshakeState(request.Nonce)
		if state == nil {
			logx.Errorf("conn(%s) HTTP set_client_DH_params: state not found", c)
			_, _ = c.Write(httpcodec.FormatResponse(nil))
			ctx.httpCodec.Reset()
			return
		}
		connId := c.ConnId()
		if err := s.pool.Submit(func() {
			dhGen := s.httpOnSetClientDHParams(state, request)
			s.eng.Trigger(connId, func(c gnet.Conn) {
				if dhGen == nil {
					_, _ = c.Write(httpcodec.FormatResponse(nil))
				} else {
					reply := func() []byte {
						xr := mtproto.GetEncodeBuf()
						defer mtproto.PutEncodeBuf(xr)
						serializeToBuffer(xr, mtproto.GenerateMessageId(), dhGen)
						return append([]byte(nil), xr.GetBuf()...)
					}()
					_, _ = c.Write(httpcodec.FormatResponse(reply))
				}
				if ctx2, _ := c.Context().(*connContext); ctx2 != nil {
					ctx2.httpCodec.Reset()
				}
			})
		}); err != nil {
			logx.Errorf("conn(%s) HTTP pool.Submit error: %v", c, err)
			_, _ = c.Write(httpcodec.FormatResponse(nil))
			ctx.httpCodec.Reset()
		}

	case *mtproto.TLMsgsAck:
		_, _ = c.Write(httpcodec.FormatResponse(nil))
		ctx.httpCodec.Reset()

	default:
		logx.Errorf("conn(%s) HTTP handshake: invalid type %T", c, obj)
		_, _ = c.Write(httpcodec.FormatResponse(nil))
		ctx.httpCodec.Reset()
	}
}

func (s *Server) onHttpEncryptedMessage(ctx *connContext, c gnet.Conn, authKeyId int64, body []byte) {
	authKey := ctx.getAuthKey()
	if authKey == nil {
		key := s.GetAuthKey(authKeyId)
		if key != nil {
			authKey = newAuthKeyUtil(key)
			ctx.putAuthKey(authKey)
		}
	} else if authKey.AuthKeyId() != authKeyId {
		logx.Errorf("conn(%s) HTTP authKeyId mismatch", c)
		_, _ = c.Write(httpcodec.FormatErrorResponse(400, "Bad Request"))
		ctx.httpCodec.Reset()
		return
	}

	if authKey != nil {
		s.doHttpEncryptedMessage(ctx, c, authKey, body)
	} else {
		bodyClone := make([]byte, len(body))
		copy(bodyClone, body)
		connId := c.ConnId()

		if err := s.pool.Submit(func() {
			queryCtx, span := otel.Tracer("gnetway").Start(context.Background(), "HttpSessionQueryAuthKey")
			defer span.End()
			key, err := s.svcCtx.Dao.SessionDispatcher.QueryAuthKey(queryCtx, authKeyId, &session.TLSessionQueryAuthKey{
				AuthKeyId: authKeyId,
			})
			if err != nil {
				logx.Errorf("conn(%s) HTTP sessionQueryAuthKey error: %v", c, err)
				s.eng.Trigger(connId, func(c gnet.Conn) {
					var (
						errPayload = make([]byte, 4)
						code       = int32(-404)
					)
					binary.LittleEndian.PutUint32(errPayload, uint32(code))
					_, _ = c.Write(httpcodec.FormatResponse(errPayload))
					if ctx2, _ := c.Context().(*connContext); ctx2 != nil {
						ctx2.httpCodec.Reset()
					}
				})
				return
			}
			s.PutAuthKey(key)
			ak := newAuthKeyUtil(key)

			s.eng.Trigger(connId, func(c gnet.Conn) {
				if ctx2, _ := c.Context().(*connContext); ctx2 != nil {
					ctx2.putAuthKey(ak)
					s.doHttpEncryptedMessage(ctx2, c, ak, bodyClone)
				}
			})
		}); err != nil {
			logx.Errorf("conn(%s) HTTP pool.Submit error: %v", c, err)
			_, _ = c.Write(httpcodec.FormatErrorResponse(502, "Bad Gateway"))
			ctx.httpCodec.Reset()
		}
	}
}

func (s *Server) doHttpEncryptedMessage(ctx *connContext, c gnet.Conn, authKey *authKeyUtil, body []byte) {
	mtpRwaData, err := authKey.AesIgeDecrypt(body[8:8+16], body[24:])
	if err != nil {
		logx.Errorf("conn(%s) HTTP decrypt error: %v", c, err)
		_, _ = c.Write(httpcodec.FormatErrorResponse(400, "Bad Request"))
		ctx.httpCodec.Reset()
		return
	}

	salt := int64(binary.LittleEndian.Uint64(mtpRwaData))
	sessionId := int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
	clientIp := ctx.clientIp
	connId := c.ConnId()

	payload := make([]byte, len(mtpRwaData[16:]))
	copy(payload, mtpRwaData[16:])

	if err := s.pool.Submit(func() {
		sendCtx, span := otel.Tracer("gnetway").Start(context.Background(), "HttpSessionSendHttpData")
		defer span.End()

		rV, err := s.svcCtx.Dao.SessionDispatcher.SendHttpData(sendCtx, authKey.AuthKeyId(), &session.TLSessionSendHttpDataToSession{
			Client: &session.SessionClientData{
				ServerId:  s.svcCtx.GatewayId,
				ConnType:  2, // TRANSPORT_HTTP
				AuthKeyId: authKey.AuthKeyId(),
				SessionId: sessionId,
				ClientIp:  clientIp,
				Salt:      salt,
				Payload:   payload,
			},
		})
		if err != nil {
			logx.Errorf("conn(%s) HTTP SendHttpData error: %v", c, err)
			s.eng.Trigger(connId, func(c gnet.Conn) {
				_, _ = c.Write(httpcodec.FormatErrorResponse(502, "Bad Gateway"))
				if ctx2, _ := c.Context().(*connContext); ctx2 != nil {
					ctx2.httpCodec.Reset()
				}
			})
			return
		}

		msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(rV.Payload)
		resp := func() []byte {
			x := mtproto.GetEncodeBuf()
			defer mtproto.PutEncodeBuf(x)
			x.Long(authKey.AuthKeyId())
			x.Bytes(msgKey)
			x.Bytes(mtpRawData)
			return append([]byte(nil), x.GetBuf()...)
		}()

		s.eng.Trigger(connId, func(c gnet.Conn) {
			_, _ = c.Write(httpcodec.FormatResponse(resp))
			if ctx2, _ := c.Context().(*connContext); ctx2 != nil {
				ctx2.httpCodec.Reset()
			}
		})
	}); err != nil {
		logx.Errorf("conn(%s) HTTP pool.Submit error: %v", c, err)
		_, _ = c.Write(httpcodec.FormatErrorResponse(502, "Bad Gateway"))
		ctx.httpCodec.Reset()
	}
}

// httpOnReqDHParams processes req_DH_params synchronously (for HTTP transport).
// This is the same logic as onReqDHParams but without asyncRun.
func (s *Server) httpOnReqDHParams(ctx *HandshakeStateCtx, request *mtproto.TLReq_DHParams) (*mtproto.Server_DH_Params, error) {
	if !bytes.Equal(request.Nonce, ctx.Nonce) {
		return nil, fmt.Errorf("invalid Nonce")
	}
	if !bytes.Equal(request.ServerNonce, ctx.ServerNonce) {
		return nil, fmt.Errorf("wrong ServerNonce")
	}
	if !bytes.Equal([]byte(request.P), p) {
		return nil, fmt.Errorf("invalid p value")
	}
	if !bytes.Equal([]byte(request.Q), q) {
		return nil, fmt.Errorf("invalid q value")
	}

	rsa := s.handshake.getKey(request.PublicKeyFingerprint)
	if rsa == nil {
		return nil, fmt.Errorf("invalid PublicKeyFingerprint")
	}

	if len(request.EncryptedData) < 256 {
		return nil, fmt.Errorf("encryptedData too short")
	}

	innerData := rsa.Decrypt([]byte(request.EncryptedData))
	if len(innerData) != 256 {
		return nil, fmt.Errorf("decrypted data length != 256")
	}

	key := innerData[:32]
	hash := crypto.Sha256Digest(innerData[32:])
	for i := 0; i < 32; i++ {
		key[i] = key[i] ^ hash[i]
	}

	paddedDataWithHash, err := crypto.NewAES256IGECryptor(key, zeroIV).Decrypt(innerData[32:])
	if err != nil {
		return nil, fmt.Errorf("AES decrypt error: %w", err)
	}

	for i, j := 0, 191; i < j; i, j = i+1, j-1 {
		paddedDataWithHash[i], paddedDataWithHash[j] = paddedDataWithHash[j], paddedDataWithHash[i]
	}

	dbuf := mtproto.GetDecodeBuf(paddedDataWithHash)
	o := dbuf.Object()
	if dbuf.GetError() != nil {
		mtproto.PutDecodeBuf(dbuf)
		return nil, fmt.Errorf("decode P_Q_inner_data error")
	}
	mtproto.PutDecodeBuf(dbuf)

	var (
		pqInnerData   *mtproto.P_QInnerData
		handshakeType int
		expiresIn     int32
	)

	switch data := o.(type) {
	case *mtproto.TLPQInnerData:
		handshakeType = mtproto.AuthKeyTypePerm
		pqInnerData = data.To_P_QInnerData()
	case *mtproto.TLPQInnerDataDc:
		handshakeType = mtproto.AuthKeyTypePerm
		pqInnerData = data.To_P_QInnerData()
	case *mtproto.TLPQInnerDataTemp:
		handshakeType = mtproto.AuthKeyTypeTemp
		expiresIn = data.GetExpiresIn()
		pqInnerData = data.To_P_QInnerData()
	case *mtproto.TLPQInnerDataTempDc:
		if data.GetDc() < 0 {
			handshakeType = mtproto.AuthKeyTypeMediaTemp
		} else {
			handshakeType = mtproto.AuthKeyTypeTemp
		}
		expiresIn = data.GetExpiresIn()
		pqInnerData = data.To_P_QInnerData()
	default:
		return nil, fmt.Errorf("invalid P_Q_inner_data type")
	}

	if !bytes.Equal([]byte(pqInnerData.GetPq()), []byte(pq)) ||
		!bytes.Equal([]byte(pqInnerData.GetP()), p) ||
		!bytes.Equal([]byte(pqInnerData.GetQ()), q) ||
		!bytes.Equal(pqInnerData.GetNonce(), request.Nonce) ||
		!bytes.Equal(pqInnerData.GetServerNonce(), request.ServerNonce) {
		return nil, fmt.Errorf("P_Q_inner_data validation failed")
	}

	ctx.NewNonce = pqInnerData.GetNewNonce()
	ctx.HandshakeType = handshakeType
	ctx.ExpiresIn = expiresIn

	A := crypto.GenerateNonce(256)
	ctx.A = A
	ctx.P = s.handshake.dh2048p

	bigIntA := new(big.Int).SetBytes(A)
	gA := new(big.Int).Exp(gBigIntDH2048G, bigIntA, gBigIntDH2048P)

	serverDHInnerData := &mtproto.TLServer_DHInnerData{Data2: &mtproto.Server_DHInnerData{
		Nonce:       request.Nonce,
		ServerNonce: request.ServerNonce,
		G:           int32(s.handshake.dh2048g[0]),
		GA:          string(gA.Bytes()),
		DhPrime:     string(ctx.P),
		ServerTime:  int32(s.CachedNow()),
	}}

	serverDHInnerDataBuf := func() []byte {
		x := mtproto.GetEncodeBuf()
		defer mtproto.PutEncodeBuf(x)
		serverDHInnerData.Encode(x, 0)
		return append([]byte(nil), x.GetBuf()...)
	}()

	tmpAesKeyAndIV := make([]byte, 64)
	var shaBuf [64]byte
	copy(shaBuf[:], ctx.NewNonce)
	copy(shaBuf[len(ctx.NewNonce):], request.ServerNonce)
	sha1A := sha1.Sum(shaBuf[:len(ctx.NewNonce)+len(request.ServerNonce)])
	copy(shaBuf[:], request.ServerNonce)
	copy(shaBuf[len(request.ServerNonce):], ctx.NewNonce)
	sha1B := sha1.Sum(shaBuf[:len(request.ServerNonce)+len(ctx.NewNonce)])
	copy(shaBuf[:], ctx.NewNonce)
	copy(shaBuf[len(ctx.NewNonce):], ctx.NewNonce)
	sha1C := sha1.Sum(shaBuf[:len(ctx.NewNonce)+len(ctx.NewNonce)])
	copy(tmpAesKeyAndIV, sha1A[:])
	copy(tmpAesKeyAndIV[20:], sha1B[:])
	copy(tmpAesKeyAndIV[40:], sha1C[:])
	copy(tmpAesKeyAndIV[60:], ctx.NewNonce[:4])

	tmpLen := 20 + len(serverDHInnerDataBuf)
	if tmpLen%16 > 0 {
		tmpLen = (tmpLen/16 + 1) * 16
	}

	tmpEncryptedAnswer := make([]byte, tmpLen)
	sha1Tmp := sha1.Sum(serverDHInnerDataBuf)
	copy(tmpEncryptedAnswer, sha1Tmp[:])
	copy(tmpEncryptedAnswer[20:], serverDHInnerDataBuf)

	e := crypto.NewAES256IGECryptor(tmpAesKeyAndIV[:32], tmpAesKeyAndIV[32:64])
	tmpEncryptedAnswer, _ = e.Encrypt(tmpEncryptedAnswer)

	ctx.State = STATE_DH_params_res

	return mtproto.MakeTLServer_DHParamsOk(&mtproto.Server_DH_Params{
		Nonce:           request.Nonce,
		ServerNonce:     request.ServerNonce,
		EncryptedAnswer: hack.String(tmpEncryptedAnswer),
	}).To_Server_DH_Params(), nil
}

// httpOnSetClientDHParams processes set_client_DH_params synchronously (for HTTP transport).
func (s *Server) httpOnSetClientDHParams(ctx *HandshakeStateCtx, request *mtproto.TLSetClient_DHParams) *mtproto.SetClient_DHParamsAnswer {
	if !bytes.Equal(request.Nonce, ctx.Nonce) {
		logx.Errorf("HTTP set_client_DH_params: wrong Nonce")
		return nil
	}
	if !bytes.Equal(request.ServerNonce, ctx.ServerNonce) {
		logx.Errorf("HTTP set_client_DH_params: wrong ServerNonce")
		return nil
	}

	bEncryptedData := []byte(request.EncryptedData)

	tmpAesKeyAndIv := make([]byte, 64)
	var shaBuf [64]byte
	copy(shaBuf[:], ctx.NewNonce)
	copy(shaBuf[len(ctx.NewNonce):], ctx.ServerNonce)
	sha1A := sha1.Sum(shaBuf[:len(ctx.NewNonce)+len(ctx.ServerNonce)])
	copy(shaBuf[:], ctx.ServerNonce)
	copy(shaBuf[len(ctx.ServerNonce):], ctx.NewNonce)
	sha1B := sha1.Sum(shaBuf[:len(ctx.ServerNonce)+len(ctx.NewNonce)])
	copy(shaBuf[:], ctx.NewNonce)
	copy(shaBuf[len(ctx.NewNonce):], ctx.NewNonce)
	sha1C := sha1.Sum(shaBuf[:len(ctx.NewNonce)+len(ctx.NewNonce)])
	copy(tmpAesKeyAndIv, sha1A[:])
	copy(tmpAesKeyAndIv[20:], sha1B[:])
	copy(tmpAesKeyAndIv[40:], sha1C[:])
	copy(tmpAesKeyAndIv[60:], ctx.NewNonce[:4])

	d := crypto.NewAES256IGECryptor(tmpAesKeyAndIv[:32], tmpAesKeyAndIv[32:64])
	decryptedData, err := d.Decrypt(bEncryptedData)
	if err != nil {
		logx.Errorf("HTTP set_client_DH_params: AES decrypt error: %v", err)
		return nil
	}

	dBuf := mtproto.GetDecodeBuf(decryptedData[20:])
	clientDHInnerData := mtproto.MakeTLClient_DHInnerData(nil)
	clientDHInnerData.Data2.Constructor = mtproto.TLConstructor(dBuf.Int())
	if err := clientDHInnerData.Decode(dBuf); err != nil {
		mtproto.PutDecodeBuf(dBuf)
		logx.Errorf("HTTP set_client_DH_params: decode error: %v", err)
		return nil
	}
	mtproto.PutDecodeBuf(dBuf)

	if !bytes.Equal(clientDHInnerData.GetNonce(), ctx.Nonce) ||
		!bytes.Equal(clientDHInnerData.GetServerNonce(), ctx.ServerNonce) {
		logx.Errorf("HTTP set_client_DH_params: nonce mismatch in inner data")
		return nil
	}

	bigIntA := new(big.Int).SetBytes(ctx.A)
	authKeyNum := new(big.Int)
	authKeyNum.Exp(new(big.Int).SetBytes([]byte(clientDHInnerData.GetGB())), bigIntA, gBigIntDH2048P)

	authKey := make([]byte, 256)
	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())

	authKeyAuxHash := make([]byte, len(ctx.NewNonce))
	copy(authKeyAuxHash, ctx.NewNonce)
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)

	authKeyId := int64(binary.LittleEndian.Uint64(authKeyAuxHash[len(ctx.NewNonce)+1+12 : len(ctx.NewNonce)+1+12+8]))

	ctx.State = STATE_dh_gen_res

	if s.saveAuthKeyInfo(ctx, mtproto.NewAuthKeyInfo(authKeyId, authKey, ctx.HandshakeType)) {
		return mtproto.MakeTLDhGenOk(&mtproto.SetClient_DHParamsAnswer{
			Nonce:         ctx.Nonce,
			ServerNonce:   ctx.ServerNonce,
			NewNonceHash1: calcNewNonceHash(ctx.NewNonce, authKey, 0x01),
		}).To_SetClient_DHParamsAnswer()
	}

	return mtproto.MakeTLDhGenRetry(&mtproto.SetClient_DHParamsAnswer{
		Nonce:         ctx.Nonce,
		ServerNonce:   ctx.ServerNonce,
		NewNonceHash2: calcNewNonceHash(ctx.NewNonce, authKey, 0x02),
	}).To_SetClient_DHParamsAnswer()
}
