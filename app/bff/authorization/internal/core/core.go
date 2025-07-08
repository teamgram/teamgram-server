// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/code/conf"
	"github.com/teamgram/teamgram-server/v2/pkg/env2"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/phonenumber"

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

func checkPhoneNumberInvalid(phone string) (string, string, error) {
	// 3. check number
	// 3.1. empty
	if phone == "" {
		// log.Errorf("check phone_number error - empty")
		return "", "", tg.Err406PhoneNumberInvalid
	}

	phone = strings.ReplaceAll(phone, " ", "")
	if phone == "+42400" ||
		phone == "+424000" ||
		phone == "+424001" ||
		phone == "+42777" {
		return "", phone[1:], nil
	}

	// fragment
	if strings.HasPrefix(phone, "+888") {
		if len(phone) == 12 {
			// +888 0888 0080
			return "", phone[1:], nil
		} else {
			return "", "", tg.Err406PhoneNumberInvalid
		}
	} else if strings.HasPrefix(phone, "888") {
		if len(phone) == 11 {
			// +888 0888 0080
			return "", phone, nil
		} else {
			return "", "", tg.Err406PhoneNumberInvalid
		}
	}

	// 3.2. check phone_number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	// We need getRegionCode from phone_number
	pNumber, err := phonenumber.MakePhoneNumberHelper(phone, "")
	if err != nil {
		// log.Errorf("check phone_number error - %v", err)
		// err = mtproto.ErrPhoneNumberInvalid
		return "", "", tg.Err406PhoneNumberInvalid
	}

	return pNumber.GetRegionCode(), pNumber.GetNormalizeDigits(), nil
}

const (
	signInMessageTpl = `Login code: %s. Do not give this code to anyone, even if they say they are from %s!

This code can be used to log in to your %s account. We never ask it for anything else.

If you didn't request this code by trying to log in on another device, simply ignore this message.`
)

func (c *AuthorizationCore) pushSignInMessage(ctx context.Context, signInUserId int64, code string) {
	time.AfterFunc(2*time.Second, func() {
		message := &tg.TLMessage{
			Out:     true,
			Date:    int32(time.Now().Unix()),
			FromId:  tg.MakePeerUser(777000),
			PeerId:  tg.MakePeerUser(signInUserId),
			Message: fmt.Sprintf(signInMessageTpl, code, env2.MyAppName, env2.MyAppName),
			Entities: []*tg.MessageEntity{
				tg.MakeMessageEntity(&tg.TLMessageEntityBold{
					Offset: 0,
					Length: 11,
				}),
				tg.MakeMessageEntity(&tg.TLMessageEntityBold{
					Offset: 22,
					Length: 3,
				}),
			},
		}

		if len(c.svcCtx.Config.SignInMessage) > 0 {
			builder := conf.ToMessageBuildHelper(
				c.svcCtx.Config.SignInMessage,
				map[string]interface{}{
					"code":     code,
					"app_name": env2.MyAppName,
				})
			message.Message, message.Entities = tg.MakeTextAndMessageEntities(builder)
		}

		_, _ = c.svcCtx.Dao.MsgClient.MsgPushUserMessage(
			ctx,
			&msgpb.TLMsgPushUserMessage{
				UserId:    777000,
				AuthKeyId: 0,
				PeerType:  tg.PEER_USER,
				PeerId:    signInUserId,
				PushType:  1,
				Message: msgpb.MakeOutboxMessage(&msgpb.TLOutboxMessage{
					NoWebpage:    false,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      message.ToMessage(),
					ScheduleDate: nil,
				}),
			})
	})
}
