// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package server

import (
	"flag"
	auth_helper "github.com/teamgram/teamgram-server/app/service/biz/auth"
	"github.com/teamgram/teamgram-server/app/service/biz/auth/auth"
	banned_helper "github.com/teamgram/teamgram-server/app/service/biz/banned"
	"github.com/teamgram/teamgram-server/app/service/biz/banned/banned"
	"github.com/teamgram/teamgram-server/app/service/biz/biz/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	// channel_helper "github.com/teamgram/teamgram-server/app/service/biz/channel"
	// "github.com/teamgram/teamgram-server/app/service/biz/channel/channel"
	chat_helper "github.com/teamgram/teamgram-server/app/service/biz/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	code_helper "github.com/teamgram/teamgram-server/app/service/biz/code"
	"github.com/teamgram/teamgram-server/app/service/biz/code/code"
	dialog_helper "github.com/teamgram/teamgram-server/app/service/biz/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	gif_helper "github.com/teamgram/teamgram-server/app/service/biz/gif"
	"github.com/teamgram/teamgram-server/app/service/biz/gif/gif"
	message_helper "github.com/teamgram/teamgram-server/app/service/biz/message"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	report_helper "github.com/teamgram/teamgram-server/app/service/biz/report"
	"github.com/teamgram/teamgram-server/app/service/biz/report/report"
	// secretchat_helper "github.com/teamgram/teamgram-server/app/service/biz/secretchat"
	// "github.com/teamgram/teamgram-server/app/service/biz/secretchat/secretchat"
	// sticker_helper "github.com/teamgram/teamgram-server/app/service/biz/sticker"
	// "github.com/teamgram/teamgram-server/app/service/biz/sticker/sticker"
	// theme_helper "github.com/teamgram/teamgram-server/app/service/biz/theme"
	// "github.com/teamgram/teamgram-server/app/service/biz/theme/theme"
	updates_helper "github.com/teamgram/teamgram-server/app/service/biz/updates"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"
	user_helper "github.com/teamgram/teamgram-server/app/service/biz/user"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	username_helper "github.com/teamgram/teamgram-server/app/service/biz/username"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
	// wallpaper_helper "github.com/teamgram/teamgram-server/app/service/biz/wallpaper"
	// "github.com/teamgram/teamgram-server/app/service/biz/wallpaper/wallpaper"
	webpage_helper "github.com/teamgram/teamgram-server/app/service/biz/webpage"
	"github.com/teamgram/teamgram-server/app/service/biz/webpage/webpage"
	// poll_helper "github.com/teamgram/teamgram-server/app/service/poll"
	// "github.com/teamgram/teamgram-server/app/service/poll/poll"
)

var configFile = flag.String("f", "etc/biz.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	// ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	s.grpcSrv = zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// auth_helper
		auth.RegisterRPCAuthServer(
			grpcServer,
			auth_helper.New(auth_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		// auth_helper
		banned.RegisterRPCBannedServer(
			grpcServer,
			banned_helper.New(banned_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		//// channel_helper
		//channel.RegisterRPCChannelServer(
		//	grpcServer,
		//	channel_helper.New(channel_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//		Mysql:         c.Mysql,
		//		Cache:         c.Cache,
		//		MediaClient:   c.MediaClient,
		//	}))

		// chat_helper
		chat.RegisterRPCChatServer(
			grpcServer,
			chat_helper.New(chat_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				MediaClient:   c.MediaClient,
			}))

		// code_helper
		code.RegisterRPCCodeServer(
			grpcServer,
			code_helper.New(code_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				KV:            c.KV,
			}))

		// dialog_helper
		dialog.RegisterRPCDialogServer(
			grpcServer,
			dialog_helper.New(dialog_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		// gif_helper
		gif.RegisterRPCGifServer(
			grpcServer,
			gif_helper.New(gif_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		// message_helper
		message.RegisterRPCMessageServer(
			grpcServer,
			message_helper.New(message_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				PollClient:    c.PollClient,
			}))

		//// poll_helper
		//poll.RegisterRPCPollServer(
		//	grpcServer,
		//	poll_helper.New(poll_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//		Mysql:         c.Mysql,
		//		Cache:         c.Cache,
		//	}))

		// report_helper
		report.RegisterRPCReportServer(
			grpcServer,
			report_helper.New(report_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		// TODO:
		//// search_helper
		//search.RegisterRPCSearchServer(
		//	grpcServer,
		//	search_helper.New(search_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//	}))

		//// sticker_helper
		//sticker.RegisterRPCStickerServer(
		//	grpcServer,
		//	sticker_helper.New(sticker_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//		Mysql:         c.Mysql,
		//		Cache:         c.Cache,
		//		MediaClient:   c.MediaClient,
		//		IdgenClient:   c.IdgenClient,
		//	}))

		//// theme_helper
		//theme.RegisterRPCThemeServer(
		//	grpcServer,
		//	theme_helper.New(theme_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//		Mysql:         c.Mysql,
		//		Cache:         c.Cache,
		//		Media:         c.MediaClient,
		//	}))

		// updates_helper
		updates.RegisterRPCUpdatesServer(
			grpcServer,
			updates_helper.New(updates_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				KV:            c.KV,
				IdgenClient:   c.IdgenClient,
			}))

		// user_helper
		user.RegisterRPCUserServer(
			grpcServer,
			user_helper.New(user_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				MediaClient:   c.MediaClient,
			}))

		// username_helper
		username.RegisterRPCUsernameServer(
			grpcServer,
			username_helper.New(username_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		//// wallpaper_helper
		//wallpaper.RegisterRPCWallpaperServer(
		//	grpcServer,
		//	wallpaper_helper.New(wallpaper_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//		Mysql:         c.Mysql,
		//		Cache:         c.Cache,
		//		Media:         c.MediaClient,
		//	}))

		// webpage_helper
		webpage.RegisterRPCWebpageServer(
			grpcServer,
			webpage_helper.New(webpage_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
			}))

		//secretchat.RegisterRPCSecretchatServer(
		//	grpcServer,
		//	secretchat_helper.New(secretchat_helper.Config{
		//		RpcServerConf: c.RpcServerConf,
		//		Mysql:         c.Mysql,
		//		KV:            c.KV,
		//		IdgenClient:   c.IdgenClient,
		//		MediaClient:   c.MediaClient,
		//	}))
	})

	// logx.Must(err)
	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
