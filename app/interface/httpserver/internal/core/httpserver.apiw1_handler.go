// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package core

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/interface/httpserver/internal/dao"
	sessionpb "github.com/teamgram/teamgram-server/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type (
	HandshakeStateCtx = dao.HandshakeStateCtx
)

const (
	STATE_ERROR              = dao.STATE_ERROR
	STATE_CONNECTED2         = dao.STATE_CONNECTED2
	STATE_HANDSHAKE          = dao.STATE_HANDSHAKE
	STATE_pq                 = dao.STATE_pq
	STATE_pq_res             = dao.STATE_pq_res
	STATE_pq_ack             = dao.STATE_pq_ack
	STATE_DH_params          = dao.STATE_DH_params
	STATE_DH_params_res      = dao.STATE_DH_params_res
	STATE_DH_params_res_fail = dao.STATE_DH_params_res_fail
	STATE_DH_params_ack      = dao.STATE_DH_params_ack
	STATE_dh_gen             = dao.STATE_dh_gen
	STATE_dh_gen_res         = dao.STATE_dh_gen_res
	STATE_dh_gen_res_retry   = dao.STATE_dh_gen_res_retry
	STATE_dh_gen_res_fail    = dao.STATE_dh_gen_res_fail
	STATE_dh_gen_ack         = dao.STATE_dh_gen_ack
	STATE_AUTH_KEY           = dao.STATE_AUTH_KEY
)

const (
	TRANSPORT_TCP  = 1 // TCP
	TRANSPORT_HTTP = 2 // HTTP
	TRANSPORT_UDP  = 3 // UDP, TODO(@benqi): 未发现有支持UDP的客户端
)

// HttpserverApiw1
// httpserver.apiw1 auth_key_id:long session_id:long payload:bytes = Bool;
func (c *HttpserverCore) HttpserverApiw1(in *mtproto.MTPRawMessage) (*mtproto.MTPRawMessage, error) {
	if in.AuthKeyId() == 0 {
		rMsg, err := c.onUnencryptedMessage(in)
		if err != nil {
			c.Logger.Errorf("httpserver.apiw1 - error: %v", err)
			return nil, err
		}

		return rMsg, nil
	} else {
		authKey, err := c.svcCtx.Dao.GetCacheAuthKey(c.ctx, in.AuthKeyId())
		if err != nil {
			c.Logger.Errorf("httpserver.apiw1 - error: %v", err)
			return nil, err
		}

		rMsg, err := c.onEncryptedMessage(crypto.NewAuthKey(authKey.AuthKeyId, authKey.AuthKey), in)
		if err != nil {
			c.Logger.Errorf("httpserver.apiw1 - error: %v", err)
			return nil, err
		}

		return rMsg, nil
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
func (c *HttpserverCore) onUnencryptedMessage(mmsg *mtproto.MTPRawMessage) (*mtproto.MTPRawMessage, error) {
	// logx.Info("receive unencryptedRawMessage: {peer: %s, ctx: %s, mmsg: %s}", conn, ctx, mmsg)
	if len(mmsg.Payload) < 8 {
		err := fmt.Errorf("invalid data len < 8")
		logx.Error(err.Error())
		return nil, err
	}

	_, obj, err := parseFromIncomingMessage(mmsg.Payload[8:])
	if err != nil {
		err := fmt.Errorf("invalid data len < 8")
		logx.Errorf(err.Error())
		return nil, err
	}

	x := mtproto.NewEncodeBuf(512)

	switch request := obj.(type) {
	case *mtproto.TLReqPq:
		logx.Infof("TLReqPq - {\"request\":%s", request)
		resPQ, err := c.svcCtx.Dao.Handshake.OnReqPq(request)
		if err != nil {
			// logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, ctx, mmsg)
			// conn.Close()
			return nil, err
		}
		c.svcCtx.Dao.PutHandshakeStateCtx(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		serializeToBuffer(x, mtproto.GenerateMessageId(), resPQ)
	case *mtproto.TLReqPqMulti:
		logx.Infof("TLReqPqMulti - {\"request\":%s", request)
		resPQ, err := c.svcCtx.Dao.Handshake.OnReqPqMulti(request)
		if err != nil {
			// logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
			// conn.Close()
			return nil, err
		}
		c.svcCtx.Dao.PutHandshakeStateCtx(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		serializeToBuffer(x, mtproto.GenerateMessageId(), resPQ)
	case *mtproto.TLReq_DHParams:
		logx.Infof("TLReq_DHParams - {\"request\":%s", request)
		if state := c.svcCtx.Dao.GetHandshakeStateCtx(request.Nonce); state != nil {
			resServerDHParam, err := c.svcCtx.Dao.Handshake.OnReqDHParams(state, obj.(*mtproto.TLReq_DHParams))
			if err != nil {
				// logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
				// conn.Close()
				return nil, err
			}
			state.State = STATE_DH_params_res
			serializeToBuffer(x, mtproto.GenerateMessageId(), resServerDHParam)
		} else {
			// logx.Errorf("onHandshake error: {invalid nonce} - {peer: %s, ctx: %s, mmsg: %s}", conn, ctx, mmsg)
			// return conn.Close()
		}
	case *mtproto.TLSetClient_DHParams:
		logx.Infof("TLSetClient_DHParams - {\"request\":%s", request)
		if state := c.svcCtx.Dao.GetHandshakeStateCtx(request.Nonce); state != nil {
			resSetClientDHParamsAnswer, err := c.svcCtx.Dao.Handshake.OnSetClientDHParams(state, obj.(*mtproto.TLSetClient_DHParams))
			if err != nil {
				// logx.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx, mmsg)
				// return conn.Close()
			}
			state.State = STATE_dh_gen_res
			serializeToBuffer(x, mtproto.GenerateMessageId(), resSetClientDHParamsAnswer)
		} else {
			// logx.Errorf("onHandshake error: {invalid nonce} - {peer: %s, ctx: %s, mmsg: %s}", conn, ctx, mmsg)
			// return conn.Close()
		}
	case *mtproto.TLMsgsAck:
		logx.Infof("TLMsgsAck - {\"request\":%s", request)
		//err = s.onMsgsAck(state, obj.(*mtproto.TLMsgsAck))
		//return nil, err
		return nil, err
	default:
		err = fmt.Errorf("invalid handshake type")
		// return conn.Close()
	}

	return &mtproto.MTPRawMessage{Payload: x.GetBuf()}, nil
}

func (c *HttpserverCore) onEncryptedMessage(authKey *crypto.AuthKey, mmsg *mtproto.MTPRawMessage) (*mtproto.MTPRawMessage, error) {
	var (
		authKeyId = mmsg.AuthKeyId()
	)

	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg.Payload[8:8+16], mmsg.Payload[24:])
	if err != nil {
		c.Logger.Errorf("onEncryptedMessage - error: %v", err)
		return nil, err
	}

	sessionId := int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))

	sessClient, err := c.svcCtx.Dao.GetSessionClient(strconv.FormatInt(mmsg.AuthKeyId(), 10))
	if err != nil {
		c.Logger.Errorf("onEncryptedMessage - error: %v", err)
		return nil, err
	}

	rV, err := sessClient.SessionSendHttpDataToSession(c.ctx, &sessionpb.TLSessionSendHttpDataToSession{
		Client: &sessionpb.SessionClientData{
			ServerId:  c.svcCtx.Dao.GetGatewayId(),
			ConnType:  TRANSPORT_HTTP,
			AuthKeyId: authKeyId,
			SessionId: sessionId,
			ClientIp:  c.MD.ClientAddr,
			QuickAck:  0,
			Salt:      int64(binary.LittleEndian.Uint64(mtpRwaData)),
			Payload:   mtpRwaData[16:],
		},
	})
	_ = rV
	if err != nil {
		c.Logger.Errorf("onEncryptedMessage - error: %v", err)
		return nil, err
	}

	msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(rV.Payload)
	x := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
	x.Long(authKey.AuthKeyId())
	x.Bytes(msgKey)
	x.Bytes(mtpRawData)
	//logx.Info("egate receiveData - ready sendToClient to: {peer: %s, auth_key_id = %d, session_id = %d}",
	//	conn,
	//	r.AuthKeyId,
	//	r.SessionId)

	return &mtproto.MTPRawMessage{Payload: x.GetBuf()}, nil
}
