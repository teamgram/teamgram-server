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

// ReportMessagesReportSpam
// report.messagesReportSpam reporter:int peer_type:int peer_id:int = Bool;
func (c *ReportCore) ReportMessagesReportSpam(in *report.TLReportMessagesReportSpam) (*mtproto.Bool, error) {
	c.svcCtx.Dao.Report(c.ctx,
		in.Reporter,
		report.MESSAGES_reportSpam,
		in.PeerType,
		in.PeerId,
		0,
		0,
		0,
		int32(report.REASON_NONE),
		"")

	return mtproto.BoolTrue, nil
}
