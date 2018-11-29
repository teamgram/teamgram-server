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

package base

import (
	"fmt"
	"github.com/nebula-chat/chatengine/mtproto"
)

type ReportReasonType int8

const (
	//inputReportReasonOther#e1746d0a text:string = ReportReason;
	REASON_OTHER ReportReasonType = 0 // 其它

	//inputReportReasonSpam#58dbcab8 = ReportReason;
	REASON_SPAM ReportReasonType = 1 // 垃圾

	//inputReportReasonViolence#1e22c78d = ReportReason;
	REASON_VIOLENCE ReportReasonType = 2 // 暴力

	//inputReportReasonPornography#2e59d922 = ReportReason;
	REASON_PORNOGRAPHY ReportReasonType = 3 // 色情
)

func (i ReportReasonType) String() (s string) {
	switch i {
	case REASON_OTHER:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	case REASON_SPAM:
		s = "nputReportReasonSpam#58dbcab8 = ReportReason"
	case REASON_VIOLENCE:
		s = "inputReportReasonPornography#2e59d922 = ReportReason"
	case REASON_PORNOGRAPHY:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	}
	return
}

func FromReportReason(reason *mtproto.ReportReason) (i ReportReasonType) {
	switch reason.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputReportReasonSpam:
		i = REASON_OTHER
	case mtproto.TLConstructor_CRC32_inputReportReasonViolence:
		i = REASON_SPAM
	case mtproto.TLConstructor_CRC32_inputReportReasonPornography:
		i = REASON_VIOLENCE
	case mtproto.TLConstructor_CRC32_inputReportReasonOther:
		i = REASON_PORNOGRAPHY
	default:
		panic(fmt.Sprintf("FromReportReason(%v) error!", reason))
	}

	return
}

func (i ReportReasonType) ToReportReason(text string) (reason *mtproto.ReportReason) {
	switch i {
	case REASON_OTHER:
		reason = &mtproto.ReportReason{
			Constructor: mtproto.TLConstructor_CRC32_inputReportReasonOther,
			Data2:       &mtproto.ReportReason_Data{Text: text},
		}
	case REASON_SPAM:
		reason = &mtproto.ReportReason{
			Constructor: mtproto.TLConstructor_CRC32_inputReportReasonSpam,
			Data2:       &mtproto.ReportReason_Data{},
		}
	case REASON_VIOLENCE:
		reason = &mtproto.ReportReason{
			Constructor: mtproto.TLConstructor_CRC32_inputReportReasonViolence,
			Data2:       &mtproto.ReportReason_Data{},
		}
	case REASON_PORNOGRAPHY:
		reason = &mtproto.ReportReason{
			Constructor: mtproto.TLConstructor_CRC32_inputReportReasonPornography,
			Data2:       &mtproto.ReportReason_Data{},
		}
	default:
		panic(fmt.Sprintf("ToReportReason(%v) error!", i))
	}
	return
}
