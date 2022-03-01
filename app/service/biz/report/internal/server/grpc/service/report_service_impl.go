/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/report/internal/core"
	"github.com/teamgram/teamgram-server/app/service/biz/report/report"
)

// ReportAccountReportPeer
// report.accountReportPeer reporter:long peer_type:int peer_id:long reason:ReportReason message:string = Bool;
func (s *Service) ReportAccountReportPeer(ctx context.Context, request *report.TLReportAccountReportPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("report.accountReportPeer - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ReportAccountReportPeer(request)
	if err != nil {
		return nil, err
	}

	c.Infof("report.accountReportPeer - reply: %s", r.DebugString())
	return r, err
}

// ReportAccountReportProfilePhoto
// report.accountReportProfilePhoto reporter:long peer_type:int peer_id:long photo_id:long reason:ReportReason message:string = Bool;
func (s *Service) ReportAccountReportProfilePhoto(ctx context.Context, request *report.TLReportAccountReportProfilePhoto) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("report.accountReportProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ReportAccountReportProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("report.accountReportProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// ReportMessagesReportSpam
// report.messagesReportSpam reporter:long peer_type:int peer_id:long = Bool;
func (s *Service) ReportMessagesReportSpam(ctx context.Context, request *report.TLReportMessagesReportSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("report.messagesReportSpam - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ReportMessagesReportSpam(request)
	if err != nil {
		return nil, err
	}

	c.Infof("report.messagesReportSpam - reply: %s", r.DebugString())
	return r, err
}

// ReportMessagesReport
// report.messagesReport reporter:long peer_type:int peer_id:long id:Vector<int> reason:ReportReason message:string = Bool;
func (s *Service) ReportMessagesReport(ctx context.Context, request *report.TLReportMessagesReport) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("report.messagesReport - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ReportMessagesReport(request)
	if err != nil {
		return nil, err
	}

	c.Infof("report.messagesReport - reply: %s", r.DebugString())
	return r, err
}

// ReportMessagesReportEncryptedSpam
// report.messagesReportEncryptedSpam reporter:long chat_id:int = Bool;
func (s *Service) ReportMessagesReportEncryptedSpam(ctx context.Context, request *report.TLReportMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("report.messagesReportEncryptedSpam - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ReportMessagesReportEncryptedSpam(request)
	if err != nil {
		return nil, err
	}

	c.Infof("report.messagesReportEncryptedSpam - reply: %s", r.DebugString())
	return r, err
}

// ReportChannelsReportSpam
// report.channelsReportSpam reporter:long channel_id:long user_id:long id:Vector<long> = Bool;
func (s *Service) ReportChannelsReportSpam(ctx context.Context, request *report.TLReportChannelsReportSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("report.channelsReportSpam - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ReportChannelsReportSpam(request)
	if err != nil {
		return nil, err
	}

	c.Infof("report.channelsReportSpam - reply: %s", r.DebugString())
	return r, err
}
