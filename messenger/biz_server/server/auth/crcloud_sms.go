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
	"time"
	//"fmt"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/http_client"
	"nebula.chat/enterprise/pkg/log"
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
	crCloudAppKey  = "4erySvd5PIIdWTZD"
	crCloudSecret  = "00S3YA5oYWUcsnQwWfzkOKygmbI4IUo3"


	intlApiUrl 	   = "http://api.1cloudsp.com/intl/api/v2/send"
	intlSignId 	   = "1410"
    intlTemplateId = "2101"

	singleApiUrl   = "http://api.1cloudsp.com/api/v2/single_send"
	singleSignId   =  "32859"
    templateId     =  "50451"

)



//var authorization string

func init() {
	//authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(jPushAppKey+":"+jPushMasterSecret))
}


type crCloudSmsResponse struct {
	Code string    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	SMUuid string  `json:"smUuid,omitempty"`
}

func SendCRCloudVerifyCode(phone string,code string) (string, error) {
	go func() {
		httpRequest := http_client.NewBeegoRequest(singleApiUrl, "POST").
			SetTimeout(time.Second*20, time.Second*10).
			Header("Connection", "close")

		//new NameValuePair("accesskey", accesskey),
		//	new NameValuePair("secret", accessSecret),
		//	new NameValuePair("sign", intlSignId),
		//	new NameValuePair("templateId", intlTemplateId),
		//	new NameValuePair("mobile", "8618916791325"),
		//	new NameValuePair("content", URLEncoder.encode("12345", "utf-8"))
		//增加正则匹配号码是中国号码，如果是走国内通道，如果不是走海外通道


		//国际
		//params := map[string]string{
		//	"accesskey":crCloudAppKey,
		//	"secret":crCloudSecret,
		//	"sign":intlSignId,
		//	"templateId": intlTemplateId,
		//	"mobile": phone,
		//	"content": code,
		//}

		//国内
		params := map[string]string{
			"accesskey":crCloudAppKey,
			"secret":crCloudSecret,
			"sign":singleSignId,
			"templateId": templateId,
			"mobile": "18916791325",
			"content": code,
		}

		httpRequest.JSONBody(params)

		r := &crCloudSmsResponse{}
		httpRequest.Debug(true)

		body, err := httpRequest.Bytes()
		glog.Infof("send_sms_help request:%s",string(httpRequest.DumpRequest()))
		glog.Infof("send_sms_help body:%s",string(body))

		if err != nil {
			log.Debugf("err - %v", err)
		}
		//if err != nil {
		//	return "", err
		//}
		//
		////err = httpRequest.ToJSON(r)
		//
		//err = json.Unmarshal(body, r)
		//if err != nil {
		//	return "", err
		//}
		//
		//if err != nil {
		//	return "", err
		//}
		//
		//return r.Code, nil
	}()

	return code, nil
}

func VerifyCRCloudCode(rCode, code string) error {
	//httpRequest := http_client.NewBeegoRequest(jPushVerifyCodeUrl+"/"+msgId+"/valid", "POST").
	//	SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
	//	SetTimeout(time.Second*20, time.Second*10).
	//	Header("Authorization", authorization)
	//
	//params := map[string]string{
	//	"code": code,
	//}
	//httpRequest.JSONBody(params)
	//
	//r := &jPushSmsResponse{}
	//httpRequest.Debug(true)
	//
	//glog.Infof("send_sms_helper request:%s",string(httpRequest.DumpRequest()))
	//body, err := httpRequest.Bytes()
	//glog.Infof("send_sms_helper request:%s",string(httpRequest.DumpRequest()))
	//glog.Infof("send_sms_helper body:%s",string(body))
	//
	////err = httpRequest.ToJSON(r)
	//if err != nil {
	//	return err
	//}
	//
	//err = json.Unmarshal(body, r)
	//if err != nil {
	//	return err
	//}
	//
	//if !r.IsValid {
	//	return fmt.Errorf("code invalid")
	//}
	return nil
}


func crCloudSendVerifyCode(phoneNumber, code, codeHash string, sentCodeType int) (string, error) {
	return SendCRCloudVerifyCode(phoneNumber,code)
}

func crCloudVerifyCode(codeHash, code, extraData string) (error) {
	return VerifyCRCloudCode(extraData, code)
}
