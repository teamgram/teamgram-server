// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package bff_proxy_client

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	idMap = map[string]string{
		"/mtproto.RPCTos":               "bff.bff", // Accepting the Terms of Service
		"/mtproto.RPCReports":           "bff.bff", // Dealing with spam and ToS violations
		"/mtproto.RPCConfiguration":     "bff.bff", // Fetching configuration
		"/mtproto.RPCQrCode":            "bff.bff", // Login via QR code
		"/mtproto.RPCMiscellaneous":     "bff.bff", // Miscellaneous
		"/mtproto.RPCAuthorization":     "bff.bff", // Registration/Authorization
		"/mtproto.RPCGdpr":              "bff.bff", // Working with GDPR export
		"/mtproto.RPCGifs":              "bff.bff", // Working with GIFs (actually MPEG4 GIFs)
		"/mtproto.RPCPromoData":         "bff.bff", // Working with Public Service Announcement and MTProxy channels
		"/mtproto.RPCTsf":               "bff.bff", // Working with TSF (internal use only)
		"/mtproto.RPCTwoFa":             "bff.bff", // Working with 2FA login
		"/mtproto.RPCSeamless":          "bff.bff", // Working with Seamless Telegram Login
		"/mtproto.RPCVoipCalls":         "bff.bff", // Working with VoIP calls
		"/mtproto.RPCChannels":          "bff.bff", // Working with channels/supergroups/geogroups
		"/mtproto.RPCChats":             "bff.bff", // Working with chats/supergroups/channels
		"/mtproto.RPCDeepLinks":         "bff.bff", // Working with deep links
		"/mtproto.RPCFiles":             "bff.bff", // Working with files
		"/mtproto.RPCWebPage":           "bff.bff", // Working with instant view pages
		"/mtproto.RPCSecretChats":       "bff.bff", // Working with secret chats
		"/mtproto.RPCPassport":          "bff.bff", // Working with telegram passport
		"/mtproto.RPCUpdates":           "bff.bff", // Working with updates
		"/mtproto.RPCInlineBot":         "bff.bff", // Working with bot inline queries and callback buttons
		"/mtproto.RPCBots":              "bff.bff", // Working with bots
		"/mtproto.RPCInternalBot":       "bff.bff", // Working with bots (internal bot API use)
		"/mtproto.RPCThemes":            "bff.bff", // Working with cloud themes
		"/mtproto.RPCContacts":          "bff.bff", // Working with contacts and top peers
		"/mtproto.RPCCreditCards":       "bff.bff", // Working with credit cards
		"/mtproto.RPCDialogs":           "bff.bff", // Working with dialogs
		"/mtproto.RPCDrafts":            "bff.bff", // Working with drafts
		"/mtproto.RPCEmoji":             "bff.bff", // Working with emoji keywords
		"/mtproto.RPCFolders":           "bff.bff", // Working with folders
		"/mtproto.RPCGames":             "bff.bff", // Working with games
		"/mtproto.RPCGroupCalls":        "bff.bff", // Working with group calls & live streaming
		"/mtproto.RPCImportedChats":     "bff.bff", // Working with imported chats
		"/mtproto.RPCLangpack":          "bff.bff", // Working with localization packs
		"/mtproto.RPCAutoDownload":      "bff.bff", // Working with media autodownload settings
		"/mtproto.RPCMessageThreads":    "bff.bff", // Working with message threads
		"/mtproto.RPCReactions":         "bff.bff", // Working with message reactions
		"/mtproto.RPCMessages":          "bff.bff", // Working with messages
		"/mtproto.RPCNotification":      "bff.bff", // Working with notification settings
		"/mtproto.RPCUsers":             "bff.bff", // Working with other users
		"/mtproto.RPCPayments":          "bff.bff", // Working with payments
		"/mtproto.RPCPolls":             "bff.bff", // Working with polls
		"/mtproto.RPCScheduledMessages": "bff.bff", // Working with scheduled messages
		"/mtproto.RPCNsfw":              "bff.bff", // Working with sensitive content (NSFW)
		"/mtproto.RPCSponsoredMessages": "bff.bff", // Working with sponsored messages
		"/mtproto.RPCProxyData":         "bff.bff", // Working with sponsored proxies
		"/mtproto.RPCStatistics":        "bff.bff", // Working with statistics
		"/mtproto.RPCStickers":          "bff.bff", // Working with stickers
		"/mtproto.RPCAccount":           "bff.bff", // Working with the user's account
		"/mtproto.RPCPhotos":            "bff.bff", // Working with user profile pictures
		"/mtproto.RPCUsernames":         "bff.bff", // Working with usernames
		"/mtproto.RPCWallpapers":        "bff.bff", // Working with wallpapers
		"/mtproto.RPCTranslate":         "bff.bff", // Working with RPCTranslate
	}
)

type BFFProxyClient struct {
	// zrpc.Client
	BFFClients map[string]zrpc.Client
}

func NewBFFProxyClient(cList []zrpc.RpcClientConf) *BFFProxyClient {
	var (
		clients   = make(map[string]zrpc.Client)
		registers = mtproto.GetRPCContextRegisters()
	)

	for _, c := range cList {
		cli := rpcx.GetCachedRpcClient(c)
		for k, v := range idMap {
			if v == c.Etcd.Key {
				clients[k] = cli
			}
		}
	}

	bizClients := make(map[string]zrpc.Client)
	for m, ctx := range registers {
		for k, _ := range idMap {
			if strings.HasPrefix(ctx.Method, k) {
				bizClients[m] = clients[k]
				break
			}
		}
	}

	return &BFFProxyClient{
		BFFClients: bizClients,
	}
}

func (c *BFFProxyClient) GetRpcClientByRequest(t interface{}) (zrpc.Client, error) {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if c2, ok := c.BFFClients[rt.Name()]; ok {
		return c2, nil
	} else {
		logx.Errorf("not found method: %s", rt.Name())
		logx.Errorf("%s blocked, License key from https://teamgram.net required to unlock enterprise features.", rt.Name())
	}

	// TODO(@benqi):
	// err := mtproto.ErrMethodNotImpl
	return nil, fmt.Errorf("not found method: %s", rt.Name())
}

// Invoke 通用grpc转发器
func (c *BFFProxyClient) Invoke(rpcMetaData *metadata.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	conn, err := c.GetRpcClientByRequest(object)
	if err != nil {
		return nil, err
	}

	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err = fmt.Errorf("Invoke error: %v not regist!\n", object)
		logx.Error(err.Error())
		return nil, mtproto.NewRpcError(status.Convert(mtproto.ErrEnterpriseIsBlocked))
	}

	// logx.Infof("Invoke - method: {%s}", t.Method)
	r := t.NewReplyFunc()
	// logx.Infof("Invoke - NewReplyFunc: {%#v}, t: {%v}", r, reflect.TypeOf(r))

	var header, trailer metadata.MD

	// Fixed @LionPuChiPuChi, 2018-12-19
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 60*time.Second)
	ctx, _ := metadata.RpcMetadataToOutgoing(ctxWithTimeout, rpcMetaData)

	rt := time.Now()

	logx.Infof("Invoke - NewReplyFunc: {%#v}", r)
	err = conn.Conn().Invoke(ctx, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))

	// log.Debugf("rpc %s time: %d", t.Method, time.Now().Unix()-rpcMetaData.ReceiveTime)
	logx.Infof("rpc Invoke: {method: %s, metadata: %s,  result: {%s}, error: {%v}}, cost = %v",
		t.Method,
		rpcMetaData.DebugString(),
		reflect.TypeOf(r),
		err,
		time.Since(rt))

	// TODO(@benqi): process header from serverF
	// grpc.Header(&header)
	// log.Debugf("Invoke - error: {%v}", err)

	if err != nil {
		logx.Errorf("RPC method: %s,  >> %v.Invoke(_) = _, %v: %#v", t.Method, conn.Conn(), err, reflect.TypeOf(err))
		if nErr, ok := status.FromError(err); ok {
			return nil, mtproto.MakeTLRpcError(&mtproto.RpcError{
				ErrorCode:    int32(nErr.Code()),
				ErrorMessage: nErr.Message(),
			})
		} else {
			rpcErr := new(mtproto.TLRpcError)
			if err2 := jsonpb.UnmarshalString(err.Error(), rpcErr); err2 == nil {
				// log.Debugf("%v", rpcErr)
				return nil, rpcErr
			} else {
				// log.Debugf("error")
				return nil, mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
			}
		}

		//case *mysql.MySQLError:
		//if rpcErr, ok := err.(*mtproto.TLRpcError); ok {
		//	return nil, rpcErr
		//} else {
		//	return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		//}
		//
		//// TODO(@benqi): 哪些情况需要断开客户端连接
		//if s, ok := status.FromError(err); ok {
		//	//switch s.Code() {
		//	//// TODO(@benqi): Rpc error, trailer has rpc_error metadata
		//	//case codes.Unknown:
		//	//	return nil, grpc_util.RpcErrorFromMD(trailer)
		//	//}
		//	return nil, mtproto.FromGRPCStatus(s)
		//} else {
		//	return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		//}
	} else {
		// log.Debugf("Invoke - Invoke reply: {%v}\n", r)
		reply, ok := r.(mtproto.TLObject)

		if !ok {
			err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
			// log.Error(err.Error())
			return nil, mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		}

		return reply, nil
	}
}
