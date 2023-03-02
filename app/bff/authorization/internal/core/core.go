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

package core

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/pkg/code/conf"
	"github.com/teamgram/teamgram-server/pkg/env2"
	"github.com/teamgram/teamgram-server/pkg/phonenumber"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizationCore {
	return &AuthorizationCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func checkPhoneNumberInvalid(phone string) (string, error) {
	// 3. check number
	// 3.1. empty
	if phone == "" {
		// log.Errorf("check phone_number error - empty")
		return "", mtproto.ErrPhoneNumberInvalid
	}

	phone = strings.ReplaceAll(phone, " ", "")
	if phone == "+42400" ||
		phone == "+42777" {
		return phone[1:], nil
	}

	// 3.2. check phone_number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	// We need getRegionCode from phone_number
	pNumber, err := phonenumber.MakePhoneNumberHelper(phone, "")
	if err != nil {
		// log.Errorf("check phone_number error - %v", err)
		// err = mtproto.ErrPhoneNumberInvalid
		return "", mtproto.ErrPhoneNumberInvalid
	}

	return pNumber.GetNormalizeDigits(), nil
}

const (
	signInMessageTpl = `Login code: %s. Do not give this code to anyone, even if they say they are from %s!

This code can be used to log in to your %s account. We never ask it for anything else.

If you didn't request this code by trying to log in on another device, simply ignore this message.`
)

func (c *AuthorizationCore) pushSignInMessage(ctx context.Context, signInUserId int64, code string) {
	time.AfterFunc(2*time.Second, func() {
		message := mtproto.MakeTLMessage(&mtproto.Message{
			Out:     true,
			Date:    int32(time.Now().Unix()),
			FromId:  mtproto.MakePeerUser(777000),
			PeerId:  mtproto.MakeTLPeerUser(&mtproto.Peer{UserId: signInUserId}).To_Peer(),
			Message: fmt.Sprintf(signInMessageTpl, code, env2.MyAppName, env2.MyAppName),
			Entities: []*mtproto.MessageEntity{
				mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: 0,
					Length: 11,
				}).To_MessageEntity(),
				mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: 22,
					Length: 3,
				}).To_MessageEntity(),
			},
		}).To_Message()

		if len(c.svcCtx.Config.SignInMessage) > 0 {
			builder := conf.ToMessageBuildHelper(
				c.svcCtx.Config.SignInMessage,
				map[string]interface{}{
					"code":     code,
					"app_name": env2.MyAppName,
				})
			message.Message, message.Entities = mtproto.MakeTextAndMessageEntities(builder)
		}

		c.svcCtx.Dao.MsgClient.MsgPushUserMessage(
			ctx,
			&msgpb.TLMsgPushUserMessage{
				UserId:    777000,
				AuthKeyId: 0,
				PeerType:  mtproto.PEER_USER,
				PeerId:    signInUserId,
				PushType:  1,
				Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
					NoWebpage:    false,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      message,
					ScheduleDate: nil,
				}).To_OutboxMessage(),
			})
	})
}
