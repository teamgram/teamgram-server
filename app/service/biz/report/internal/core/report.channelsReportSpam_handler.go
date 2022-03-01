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

// ReportChannelsReportSpam
// report.channelsReportSpam reporter:int channel_id:int user_id:int id:Vector<int> = Bool;
func (c *ReportCore) ReportChannelsReportSpam(in *report.TLReportChannelsReportSpam) (*mtproto.Bool, error) {
	c.svcCtx.Dao.ReportIdList(c.ctx,
		in.Reporter,
		report.CHANNELS_reportSpam,
		mtproto.PEER_CHANNEL,
		in.ChannelId,
		0,
		in.UserId,
		in.Id,
		int32(report.REASON_NONE),
		"")

	return mtproto.BoolTrue, nil
}
