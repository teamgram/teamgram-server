/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/reports/internal/core"
)

// AccountReportPeer
// account.reportPeer#c5ba3d86 peer:InputPeer reason:ReportReason message:string = Bool;
func (s *Service) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.reportPeer - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountReportPeer(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.reportPeer - reply: %s", r.DebugString())
	return r, err
}

// AccountReportProfilePhoto
// account.reportProfilePhoto#fa8cc6f5 peer:InputPeer photo_id:InputPhoto reason:ReportReason message:string = Bool;
func (s *Service) AccountReportProfilePhoto(ctx context.Context, request *mtproto.TLAccountReportProfilePhoto) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.reportProfilePhoto - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountReportProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.reportProfilePhoto - reply: %s", r.DebugString())
	return r, err
}

// MessagesReportSpam
// messages.reportSpam#cf1592db peer:InputPeer = Bool;
func (s *Service) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.reportSpam - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReportSpam(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.reportSpam - reply: %s", r.DebugString())
	return r, err
}

// MessagesReport
// messages.report#8953ab4e peer:InputPeer id:Vector<int> reason:ReportReason message:string = Bool;
func (s *Service) MessagesReport(ctx context.Context, request *mtproto.TLMessagesReport) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.report - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReport(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.report - reply: %s", r.DebugString())
	return r, err
}

// MessagesReportEncryptedSpam
// messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;
func (s *Service) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.reportEncryptedSpam - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReportEncryptedSpam(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.reportEncryptedSpam - reply: %s", r.DebugString())
	return r, err
}

// ChannelsReportSpam
// channels.reportSpam#f44a8315 channel:InputChannel participant:InputPeer id:Vector<int> = Bool;
func (s *Service) ChannelsReportSpam(ctx context.Context, request *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("channels.reportSpam - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ChannelsReportSpam(request)
	if err != nil {
		return nil, err
	}

	c.Infof("channels.reportSpam - reply: %s", r.DebugString())
	return r, err
}
