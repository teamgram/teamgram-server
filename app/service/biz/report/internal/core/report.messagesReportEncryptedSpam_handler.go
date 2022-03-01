/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/report/report"
)

// ReportMessagesReportEncryptedSpam
// report.messagesReportEncryptedSpam reporter:int chat_id:int = Bool;
func (c *ReportCore) ReportMessagesReportEncryptedSpam(in *report.TLReportMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	c.svcCtx.Dao.Report(c.ctx,
		in.Reporter,
		report.MESSAGES_reportEncryptedSpam,
		mtproto.PEER_ENCRYPTED_CHAT,
		int64(in.ChatId),
		0,
		0,
		0,
		int32(report.REASON_NONE),
		"")

	return mtproto.BoolTrue, nil
}
