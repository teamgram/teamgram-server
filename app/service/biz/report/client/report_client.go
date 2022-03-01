/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package report_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/report/report"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ReportClient interface {
	ReportAccountReportPeer(ctx context.Context, in *report.TLReportAccountReportPeer) (*mtproto.Bool, error)
	ReportAccountReportProfilePhoto(ctx context.Context, in *report.TLReportAccountReportProfilePhoto) (*mtproto.Bool, error)
	ReportMessagesReportSpam(ctx context.Context, in *report.TLReportMessagesReportSpam) (*mtproto.Bool, error)
	ReportMessagesReport(ctx context.Context, in *report.TLReportMessagesReport) (*mtproto.Bool, error)
	ReportMessagesReportEncryptedSpam(ctx context.Context, in *report.TLReportMessagesReportEncryptedSpam) (*mtproto.Bool, error)
	ReportChannelsReportSpam(ctx context.Context, in *report.TLReportChannelsReportSpam) (*mtproto.Bool, error)
}

type defaultReportClient struct {
	cli zrpc.Client
}

func NewReportClient(cli zrpc.Client) ReportClient {
	return &defaultReportClient{
		cli: cli,
	}
}

// ReportAccountReportPeer
// report.accountReportPeer reporter:long peer_type:int peer_id:long reason:ReportReason message:string = Bool;
func (m *defaultReportClient) ReportAccountReportPeer(ctx context.Context, in *report.TLReportAccountReportPeer) (*mtproto.Bool, error) {
	client := report.NewRPCReportClient(m.cli.Conn())
	return client.ReportAccountReportPeer(ctx, in)
}

// ReportAccountReportProfilePhoto
// report.accountReportProfilePhoto reporter:long peer_type:int peer_id:long photo_id:long reason:ReportReason message:string = Bool;
func (m *defaultReportClient) ReportAccountReportProfilePhoto(ctx context.Context, in *report.TLReportAccountReportProfilePhoto) (*mtproto.Bool, error) {
	client := report.NewRPCReportClient(m.cli.Conn())
	return client.ReportAccountReportProfilePhoto(ctx, in)
}

// ReportMessagesReportSpam
// report.messagesReportSpam reporter:long peer_type:int peer_id:long = Bool;
func (m *defaultReportClient) ReportMessagesReportSpam(ctx context.Context, in *report.TLReportMessagesReportSpam) (*mtproto.Bool, error) {
	client := report.NewRPCReportClient(m.cli.Conn())
	return client.ReportMessagesReportSpam(ctx, in)
}

// ReportMessagesReport
// report.messagesReport reporter:long peer_type:int peer_id:long id:Vector<int> reason:ReportReason message:string = Bool;
func (m *defaultReportClient) ReportMessagesReport(ctx context.Context, in *report.TLReportMessagesReport) (*mtproto.Bool, error) {
	client := report.NewRPCReportClient(m.cli.Conn())
	return client.ReportMessagesReport(ctx, in)
}

// ReportMessagesReportEncryptedSpam
// report.messagesReportEncryptedSpam reporter:long chat_id:int = Bool;
func (m *defaultReportClient) ReportMessagesReportEncryptedSpam(ctx context.Context, in *report.TLReportMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	client := report.NewRPCReportClient(m.cli.Conn())
	return client.ReportMessagesReportEncryptedSpam(ctx, in)
}

// ReportChannelsReportSpam
// report.channelsReportSpam reporter:long channel_id:long user_id:long id:Vector<long> = Bool;
func (m *defaultReportClient) ReportChannelsReportSpam(ctx context.Context, in *report.TLReportChannelsReportSpam) (*mtproto.Bool, error) {
	client := report.NewRPCReportClient(m.cli.Conn())
	return client.ReportChannelsReportSpam(ctx, in)
}
