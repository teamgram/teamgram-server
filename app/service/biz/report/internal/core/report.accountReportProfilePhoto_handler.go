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

// ReportAccountReportProfilePhoto
// report.accountReportProfilePhoto reporter:int peer_type:int peer_id:int photo_id:long reason:ReportReason message:string = Bool;
func (c *ReportCore) ReportAccountReportProfilePhoto(in *report.TLReportAccountReportProfilePhoto) (*mtproto.Bool, error) {
	reason := report.FromReportReason(in.GetReason())

	c.svcCtx.Dao.Report(c.ctx,
		in.Reporter,
		report.ACCOUNTS_reportProfilePhoto,
		in.PeerType,
		in.PeerId,
		in.PhotoId,
		0,
		0,
		int32(reason),
		in.Message)

	return mtproto.BoolTrue, nil
}
