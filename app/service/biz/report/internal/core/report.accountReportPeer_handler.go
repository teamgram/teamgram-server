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

// ReportAccountReportPeer
// report.accountReportPeer reporter:int peer_type:int peer_id:int reason:ReportReason message:string = Bool;
func (c *ReportCore) ReportAccountReportPeer(in *report.TLReportAccountReportPeer) (*mtproto.Bool, error) {
	reason := report.FromReportReason(in.GetReason())

	c.svcCtx.Dao.Report(c.ctx,
		in.Reporter,
		report.ACCOUNTS_reportPeer,
		in.PeerType,
		in.PeerId,
		0,
		0,
		0,
		int32(reason),
		in.Message)

	return mtproto.BoolTrue, nil
}
