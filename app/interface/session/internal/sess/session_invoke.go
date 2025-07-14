// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package sess

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/mt"
	"github.com/teamgram/proto/v2/tg"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
)

func (c *session) onInvokeWithLayer(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithLayer) {
	logx.WithContext(ctx).Infof("onInvokeWithLayer - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeWithLayer Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	//dBuf.Object()
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.sessList.cb.onUpdateLayer(ctx, request.Layer)

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsg(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeAfterMsg) {
	logx.WithContext(ctx).Infof("onInvokeAfterMsg - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeAfterMsg Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsgs(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeAfterMsgs) {
	logx.WithContext(ctx).Infof("onInvokeAfterMsgs - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeAfterMsgs Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	/*
		if invokeAfterMsgs.GetQuery() == nil {
			logx.Errorf("invokeAfterMsgs Query is nil, query: {%v}", invokeAfterMsgs)
			return
		}

		dBuf := bin.NewDecoder(invokeAfterMsgs.Query)
		query, err := iface.DecodeObject(dBuf)
		if err != nil {
			logx.Errorf("Decode query error: %v", err)
			return
		}

		if len(invokeAfterMsgs.MsgIds) == 0 {
			// TODO(@benqi): invalid msgIds, ignore??

			messages[i].Object = query
		} else {
			var maxMsgId = invokeAfterMsgs.MsgIds[0]
			for j := 1; j < len(invokeAfterMsgs.MsgIds); j++ {
				if maxMsgId > invokeAfterMsgs.MsgIds[j] {
					maxMsgId = invokeAfterMsgs.MsgIds[j]
				}
			}


			var found = false
			for j := 0; j < i; j++ {
				if messages[j].MsgId == maxMsgId {
					messages[i].Object = query
					found = true
					break
				}
			}

			if !found {
				for j := i + 1; j < len(messages); j++ {
					if messages[j].MsgId == maxMsgId {
						// c.messages[i].Object = query
						messages[i].Object = query
						found = true
						messages = append(messages, messages[i])

						// set messages[i] = nil, will ignore this.
						messages[i] = nil
						break
					}
				}
			}

			if !found {
				// TODO(@benqi): backup message, wait.

				messages[i].Object = query
			}
		}
	*/

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithoutUpdates(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithoutUpdates) {
	logx.WithContext(ctx).Infof("onInvokeWithoutUpdates - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(request))

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeWithoutUpdates Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithMessagesRange(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithMessagesRange) {
	logx.WithContext(ctx).Infof("onInvokeWithMessagesRange - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(request))

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeWithMessagesRange Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithTakeout(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithTakeout) {
	logx.WithContext(ctx).Infof("onInvokeWithTakeout - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(request))

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeWithTakeout Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithBusinessConnection(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithBusinessConnection) {
	logx.WithContext(ctx).Infof("onInvokeWithBusinessConnection - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(request))

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeWithBusinessConnection Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithGooglePlayIntegrity(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithGooglePlayIntegrity) {
	logx.WithContext(ctx).Infof("onInvokeWithGooglePlayIntegrity - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, nonce: %s, token: %s, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request.Nonce,
		request.Token,
		reflect.TypeOf(request))

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("onInvokeWithGooglePlayIntegrity Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

// onInvokeWithApnsSecret
func (c *session) onInvokeWithApnsSecret(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInvokeWithApnsSecret) {
	logx.WithContext(ctx).Infof("onInvokeWithApnsSecret - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, nonce: %s, secret: %s, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request.Nonce,
		request.Secret,
		reflect.TypeOf(request))

	if request.Query == nil {
		logx.WithContext(ctx).Errorf("invokeWithApnsSecret Query is nil, query: {%s}", request)
		return
	}

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInitConnection(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *tg.TLInitConnection) {
	logx.WithContext(ctx).Infof("onInitConnection - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	c.sessList.cb.onUpdateInitConnection(ctx, clientIp, request)

	dBuf := bin.NewDecoder(request.Query)
	query, err := iface.DecodeObject(dBuf)
	if err != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %s", err)
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onRpcRequest(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, query iface.TLObject) bool {
	logx.WithContext(ctx).Infof("onRpcRequest - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(query))

	// TODO(@benqi): sync AuthUserId??
	//requestMessage := &mt.TLMessage2{
	//	MsgId:  msgId.msgId,
	//	Seqno:  msgId.seqNo,
	//	Object: object,
	//}

	switch query.(type) {
	case *tg.TLAccountRegisterDevice:
		registerDevice, _ := query.(*tg.TLAccountRegisterDevice)
		if registerDevice.TokenType == 7 {
			pushSessionId, _ := strconv.ParseInt(registerDevice.Token, 10, 64)
			c.sessList.cb.onBindPushSessionId(ctx, c.sessList, pushSessionId)
		}
	case *tg.TLUpdatesGetState:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
	case *tg.TLUpdatesGetDifference:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
	case *tg.TLUpdatesGetChannelDifference:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
		//case *tg.TLAuthBindTempAuthKey:
		//	res, err := c.AuthSessionRpcClient.AuthBindTempAuthKey(context.Background(), query.(*tg.TLAuthBindTempAuthKey))
		//	if err != nil {
		//		logx.WithContext(ctx).Errorf("bindTempAuthKey error - %v", err)
		//		err = tg.ErrInternalServerError
		//		c.sendRpcResultToQueue(gatewayId, msgId.msgId, &mt.RpcError{
		//			ErrorCode:    500,
		//			ErrorMessage: "INTERNAL_SERVER_ERROR",
		//		})
		//	} else {
		//		c.sendRpcResultToQueue(gatewayId, msgId.msgId, res)
		//	}
		//	msgId.state = RECEIVED | RESPONSE_GENERATED
		//	return false

		//case *tg.TLAuthSignIn:
		//	if !c.isGeneric {
		//		c.isGeneric = true
		//		c.cb.setOnline()
		//	}
		//case *tg.TLAuthSignUp:
		//	if !c.isGeneric {
		//		c.isGeneric = true
		//		c.cb.setOnline()
		//	}
	case *tg.TLAccountUpdateStatus:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
	case *tg.TLUsersGetUsers:
		// logx.WithContext(ctx).Infof("user.getUsers: %s", query)
	}

	switch c.sessList.cb.state {
	case tg.AuthStateNormal:
		// state is ok
	case tg.AuthStateNeedPassword:
		if !checkRpcWithoutLogin(query) {
			c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, tg.NewRpcError(tg.ErrSessionPasswordNeeded))
			msgId.state = RECEIVED | RESPONSE_GENERATED
			return false
		}
	default:
		if !checkRpcWithoutLogin(query) {
			c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, tg.NewRpcError(tg.ErrAuthKeyUnregistered))
			msgId.state = RECEIVED | RESPONSE_GENERATED
			return false
		}
	}

	msgId.state = RECEIVED | RPC_PROCESSING

	c.sessList.cb.tmpRpcApiMessageList = append(
		c.sessList.cb.tmpRpcApiMessageList,
		&rpcApiMessage{
			ctx:       contextx.ValueOnlyFrom(ctx),
			sessList:  c.sessList,
			sessionId: c.sessionId,
			clientIp:  clientIp,
			reqMsgId:  msgId.msgId,
			reqMsg:    query,
		})

	return true
}

func (c *session) onRpcResult(ctx context.Context, rpcResult *rpcApiMessage) {
	rpcErr, _ := rpcResult.TryGetRpcResultError()
	if rpcErr != nil && rpcErr.ErrorCode == int32(tg.ErrNotReturnClient) {
		logx.WithContext(ctx).Debugf("recv NOTRETURN_CLIENT")
		c.pendingQueue.Add(rpcResult.reqMsgId)
		return
	}

	switch request := rpcResult.reqMsg.(type) {
	case *tg.TLAuthBindTempAuthKey:
		if rpcErr != nil {
			_ = request
			if rpcErr.Message() == "ENCRYPTED_MESSAGE_INVALID" {
				c.sessList.cb.changeAuthState(ctx, tg.AuthStateUnknown, 0)
				// c.sessList.cb.cb.DeleteByAuthKeyId(c.sessList.authId)
			}
		} else {
			c.sessList.changeAuthState(tg.AuthStatePermBound)
		}
	case *tg.TLAuthLogOut:
		c.sessList.cb.changeAuthState(ctx, tg.AuthStateLogout, 0)
	case *tg.TLAuthSendCode:
		if rpcErr == nil {
			//authSentCode, _ := rpcResult.rpcResult.Result.(*tg.AuthSentCode)
			//if authSentCode.AuthSentCodeClazzName() == tg.ClazzName_auth_sentCodeSuccess {
			//	c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, authSentCode.GetAuthorization().GetUser().GetId())
			//}

			if authSentCode, ok := rpcResult.rpcResult.Result.(*tg.AuthSentCode); ok {
				authSentCode.Match(func(c2 *tg.TLAuthSentCodeSuccess) interface{} {
					c3, _ := c2.Authorization.(*tg.TLAuthAuthorization)
					c4, _ := c3.User.(*tg.TLUser)
					c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, c4.Id)

					return nil
				})
			}
		} else {
			// hack
			errMsg := rpcErr.Message()
			if strings.HasPrefix(errMsg, "SESSION_PASSWORD_NEEDED_") {
				vId := errMsg[len("SESSION_PASSWORD_NEEDED_"):]
				hackId, _ := strconv.ParseInt(vId, 10, 64)
				rpcErr.ErrorMessage = "SESSION_PASSWORD_NEEDED"
				c.sessList.cb.changeAuthState(ctx, tg.AuthStateNeedPassword, hackId)
			}
		}
	case *tg.TLAuthExportLoginToken:
		if rpcErr == nil {
			// authLoginToken, _ := rpcResult.rpcResult.Result.(*tg.AuthLoginToken)
			// if authLoginToken.GetPredicateName() == tg.ClazzName_auth_loginTokenSuccess {
			//	c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, authLoginToken.GetAuthorization().GetUser().GetId())
			// }

			if authLoginToken, ok := rpcResult.rpcResult.Result.(*tg.AuthLoginToken); ok {
				authLoginToken.Match(func(c2 *tg.TLAuthLoginTokenSuccess) interface{} {
					c3, _ := c2.Authorization.(*tg.TLAuthAuthorization)
					c4, _ := c3.User.(*tg.TLUser)
					c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, c4.Id)

					return nil
				})
			}
		} else {
			// hack
			errMsg := rpcErr.Message()
			if strings.HasPrefix(errMsg, "SESSION_PASSWORD_NEEDED_") {
				vId := errMsg[len("SESSION_PASSWORD_NEEDED_"):]
				hackId, _ := strconv.ParseInt(vId, 10, 64)
				rpcErr.ErrorMessage = "SESSION_PASSWORD_NEEDED"
				c.sessList.cb.changeAuthState(ctx, tg.AuthStateNeedPassword, hackId)
			}
		}
	case *tg.TLAuthSignIn:
		if rpcErr == nil {
			//authAuthorization, _ := rpcResult.rpcResult.Result.(*tg.AuthAuthorization)
			//if authAuthorization.GetPredicateName() == tg.ClazzName_auth_authorization {
			//	c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, authAuthorization.GetUser().GetId())
			//}

			if authAuthorization, ok := rpcResult.rpcResult.Result.(*tg.AuthAuthorization); ok {
				authAuthorization.Match(func(c2 *tg.TLAuthAuthorization) interface{} {
					c3, _ := c2.User.(*tg.TLUser)
					c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, c3.Id)

					return nil
				})
			}
		} else {
			// hack
			errMsg := rpcErr.Message()
			if strings.HasPrefix(errMsg, "SESSION_PASSWORD_NEEDED_") {
				vId := errMsg[len("SESSION_PASSWORD_NEEDED_"):]
				hackId, _ := strconv.ParseInt(vId, 10, 64)
				rpcErr.ErrorMessage = "SESSION_PASSWORD_NEEDED"
				c.sessList.cb.changeAuthState(ctx, tg.AuthStateNeedPassword, hackId)
			}
		}
	case *tg.TLAuthSignUp:
		if rpcErr == nil {
			//authAuthorization, _ := rpcResult.rpcResult.Result.(*tg.AuthAuthorization)
			//if authAuthorization.GetPredicateName() == tg.ClazzName_auth_authorization {
			//	c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, authAuthorization.GetUser().GetId())
			//}

			if authAuthorization, ok := rpcResult.rpcResult.Result.(*tg.AuthAuthorization); ok {
				authAuthorization.Match(func(c2 *tg.TLAuthAuthorization) interface{} {
					c3, _ := c2.User.(*tg.TLUser)
					c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, c3.Id)

					return nil
				})
			}
		}
	case *tg.TLAuthImportAuthorization:
		if rpcErr == nil && c.sessList.cb.state == tg.AuthStateNeedPassword {
			//authAuthorization, _ := rpcResult.rpcResult.Result.(*tg.AuthAuthorization)
			//if authAuthorization.GetPredicateName() == tg.ClazzName_auth_authorization {
			//	c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, authAuthorization.GetUser().GetId())
			//}

			if authAuthorization, ok := rpcResult.rpcResult.Result.(*tg.AuthAuthorization); ok {
				authAuthorization.Match(func(c2 *tg.TLAuthAuthorization) interface{} {
					c3, _ := c2.User.(*tg.TLUser)
					c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, c3.Id)

					return nil
				})
			}
		}
	case *tg.TLAuthCheckPassword:
		if rpcErr == nil && c.sessList.cb.state == tg.AuthStateNeedPassword {
			//authAuthorization, _ := rpcResult.rpcResult.Result.(*tg.AuthAuthorization)
			//if authAuthorization.GetPredicateName() == tg.ClazzName_auth_authorization {
			//	c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, authAuthorization.GetUser().GetId())
			//}

			if authAuthorization, ok := rpcResult.rpcResult.Result.(*tg.AuthAuthorization); ok {
				authAuthorization.Match(func(c2 *tg.TLAuthAuthorization) interface{} {
					c3, _ := c2.User.(*tg.TLUser)
					c.sessList.cb.changeAuthState(ctx, tg.AuthStateNormal, c3.Id)

					return nil
				})
			}
		}
	default:
	}

	c.sendRpcResult(ctx, rpcResult.MoveRpcResult())
}

func (c *session) sendRpcResult(ctx context.Context, rpcResult *mt.TLRpcResult) {
	// TODO(@benqi): lookup inBoxMsg
	msgId := c.inQueue.Lookup(rpcResult.ReqMsgId)
	if msgId == nil {
		logx.WithContext(ctx).Errorf("not found msgId, maybe removed: %d", rpcResult.ReqMsgId)
		return
	}

	gatewayId := c.getGatewayId()
	c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, rpcResult.Result)
	msgId.state = RECEIVED | ACKNOWLEDGED

	if gatewayId == "" {
		logx.WithContext(ctx).Errorf("gatewayId is empty, send delay...")
	} else {
		c.sendQueueToGateway(ctx, gatewayId)
	}
}
