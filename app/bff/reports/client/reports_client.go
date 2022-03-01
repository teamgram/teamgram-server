/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package reports_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ReportsClient interface {
	AccountReportPeer(ctx context.Context, in *mtproto.TLAccountReportPeer) (*mtproto.Bool, error)
	AccountReportProfilePhoto(ctx context.Context, in *mtproto.TLAccountReportProfilePhoto) (*mtproto.Bool, error)
	MessagesReportSpam(ctx context.Context, in *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error)
	MessagesReport(ctx context.Context, in *mtproto.TLMessagesReport) (*mtproto.Bool, error)
	MessagesReportEncryptedSpam(ctx context.Context, in *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error)
	ChannelsReportSpam(ctx context.Context, in *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error)
}

type defaultReportsClient struct {
	cli zrpc.Client
}

func NewReportsClient(cli zrpc.Client) ReportsClient {
	return &defaultReportsClient{
		cli: cli,
	}
}

// AccountReportPeer
// account.reportPeer#c5ba3d86 peer:InputPeer reason:ReportReason message:string = Bool;
func (m *defaultReportsClient) AccountReportPeer(ctx context.Context, in *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReportsClient(m.cli.Conn())
	return client.AccountReportPeer(ctx, in)
}

// AccountReportProfilePhoto
// account.reportProfilePhoto#fa8cc6f5 peer:InputPeer photo_id:InputPhoto reason:ReportReason message:string = Bool;
func (m *defaultReportsClient) AccountReportProfilePhoto(ctx context.Context, in *mtproto.TLAccountReportProfilePhoto) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReportsClient(m.cli.Conn())
	return client.AccountReportProfilePhoto(ctx, in)
}

// MessagesReportSpam
// messages.reportSpam#cf1592db peer:InputPeer = Bool;
func (m *defaultReportsClient) MessagesReportSpam(ctx context.Context, in *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReportsClient(m.cli.Conn())
	return client.MessagesReportSpam(ctx, in)
}

// MessagesReport
// messages.report#8953ab4e peer:InputPeer id:Vector<int> reason:ReportReason message:string = Bool;
func (m *defaultReportsClient) MessagesReport(ctx context.Context, in *mtproto.TLMessagesReport) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReportsClient(m.cli.Conn())
	return client.MessagesReport(ctx, in)
}

// MessagesReportEncryptedSpam
// messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;
func (m *defaultReportsClient) MessagesReportEncryptedSpam(ctx context.Context, in *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReportsClient(m.cli.Conn())
	return client.MessagesReportEncryptedSpam(ctx, in)
}

// ChannelsReportSpam
// channels.reportSpam#f44a8315 channel:InputChannel participant:InputPeer id:Vector<int> = Bool;
func (m *defaultReportsClient) ChannelsReportSpam(ctx context.Context, in *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	client := mtproto.NewRPCReportsClient(m.cli.Conn())
	return client.ChannelsReportSpam(ctx, in)
}
