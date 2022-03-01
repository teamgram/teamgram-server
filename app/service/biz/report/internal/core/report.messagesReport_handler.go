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

// ReportMessagesReport
// report.messagesReport reporter:int peer_type:int peer_id:int id:Vector<int> reason:ReportReason message:string = Bool;
func (c *ReportCore) ReportMessagesReport(in *report.TLReportMessagesReport) (*mtproto.Bool, error) {
	reason := report.FromReportReason(in.Reason)

	c.svcCtx.Dao.ReportIdList(c.ctx,
		in.Reporter,
		report.MESSAGES_report,
		in.PeerType,
		in.PeerId,
		0,
		0,
		in.Id,
		int32(reason),
		in.Message)

	return mtproto.BoolTrue, nil
}
