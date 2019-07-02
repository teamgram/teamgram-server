// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package auth

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/http_client"
	"time"
)

// TODO(@benqi): impl sendSms

//type SendSmsVerifyCodeF func(phoneNumber, code, codeHash string, sentCodeType int) (string, error)
//type VerifySmsCodeF func(codeHash, code, extraData string) (error)
//
//func getSendSmsVerifyCodeF() SendSmsVerifyCodeF {
//	return jSendVerifyCode
//}
//
//func getVerifySmsCodeF() VerifySmsCodeF {
//	return jVerifyCode
//}

const (
	jPushAppKey                = "cde81b2e17ca49f16fb226f1"
	jPushMasterSecret          = "12cbe11ab020e15d4b4a7042"
	jPushSignatureId           = "9257"
	jPushLoginVerifyTemplateId = "1" // 手机验证码短信模板
)

const (
	jPushVerifyCodeUrl = "https://api.sms.jpush.cn/v1/codes"
)

var authorization string

func init() {
	authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(jPushAppKey+":"+jPushMasterSecret))
}

type error2 struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type jPushSmsResponse struct {
	MsgId   string `json:"msg_id,omitempty"`
	IsValid bool   `json:"is_valid,omitempty"`
	Error   error2 `json:"error,omitempty"`
}

func SendJPushVerifyCode(phone string) (string, error) {
	httpRequest := http_client.NewBeegoRequest(jPushVerifyCodeUrl, "POST").
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTimeout(time.Second*20, time.Second*10).
		Header("Authorization", authorization)

	params := map[string]string{
		"mobile": phone,
		"temp_id": jPushLoginVerifyTemplateId,
		"sign_id": jPushSignatureId,
	}
	httpRequest.JSONBody(params)

	r := &jPushSmsResponse{}
	httpRequest.Debug(true)

	body, err := httpRequest.Bytes()
	glog.Infof("send_sms_help request:%s",string(httpRequest.DumpRequest()))
	glog.Infof("send_sms_help body:%s",string(body))

	if err != nil {
		return "", err
	}

	// err := httpRequest.ToJSON(r)

	err = json.Unmarshal(body, r)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return r.MsgId, nil
}

func VerifyJPushCode(msgId, code string) error {
	httpRequest := http_client.NewBeegoRequest(jPushVerifyCodeUrl+"/"+msgId+"/valid", "POST").
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTimeout(time.Second*20, time.Second*10).
		Header("Authorization", authorization)

	params := map[string]string{
		"code": code,
	}
	httpRequest.JSONBody(params)

	r := &jPushSmsResponse{}
	httpRequest.Debug(true)

	glog.Infof("send_sms_helper request:%s",string(httpRequest.DumpRequest()))
	body, err := httpRequest.Bytes()
	glog.Infof("send_sms_helper request:%s",string(httpRequest.DumpRequest()))
	glog.Infof("send_sms_helper body:%s",string(body))

	//err = httpRequest.ToJSON(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}

	if !r.IsValid {
		return fmt.Errorf("code invalid")
	}
	return nil
}


func jSendVerifyCode(phoneNumber, code, codeHash string, sentCodeType int) (string, error) {
	return SendJPushVerifyCode(phoneNumber)
}

func jVerifyCode(codeHash, code, extraData string) (error) {
	return VerifyJPushCode(extraData, code)
}
