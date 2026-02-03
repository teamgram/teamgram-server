// Copyright 2022 Teamgooo Authors
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

	accounthelper "github.com/teamgooo/teamgooo-server/app/bff/account"
	"github.com/teamgooo/teamgooo-server/app/bff/account/account/accountservice"
	authorizationhelper "github.com/teamgooo/teamgooo-server/app/bff/authorization"
	"github.com/teamgooo/teamgooo-server/app/bff/authorization/authorization/authorizationservice"
	autodownloadhelper "github.com/teamgooo/teamgooo-server/app/bff/autodownload"
	"github.com/teamgooo/teamgooo-server/app/bff/autodownload/autodownload/autodownloadservice"
	"github.com/teamgooo/teamgooo-server/app/bff/bff/internal/config"
	chatinviteshelper "github.com/teamgooo/teamgooo-server/app/bff/chatinvites"
	"github.com/teamgooo/teamgooo-server/app/bff/chatinvites/chatinvites/chatinvitesservice"
	chatshelper "github.com/teamgooo/teamgooo-server/app/bff/chats"
	"github.com/teamgooo/teamgooo-server/app/bff/chats/chats/chatsservice"
	configurationhelper "github.com/teamgooo/teamgooo-server/app/bff/configuration"
	"github.com/teamgooo/teamgooo-server/app/bff/configuration/configuration/configurationservice"
	contactshelper "github.com/teamgooo/teamgooo-server/app/bff/contacts"
	"github.com/teamgooo/teamgooo-server/app/bff/contacts/contacts/contactsservice"
	dialogshelper "github.com/teamgooo/teamgooo-server/app/bff/dialogs"
	"github.com/teamgooo/teamgooo-server/app/bff/dialogs/dialogs/dialogsservice"
	draftshelper "github.com/teamgooo/teamgooo-server/app/bff/drafts"
	"github.com/teamgooo/teamgooo-server/app/bff/drafts/drafts/draftsservice"
	fileshelper "github.com/teamgooo/teamgooo-server/app/bff/files"
	"github.com/teamgooo/teamgooo-server/app/bff/files/files/filesservice"
	messageshelper "github.com/teamgooo/teamgooo-server/app/bff/messages"
	"github.com/teamgooo/teamgooo-server/app/bff/messages/messages/messagesservice"
	miscellaneoushelper "github.com/teamgooo/teamgooo-server/app/bff/miscellaneous"
	"github.com/teamgooo/teamgooo-server/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
	notificationhelper "github.com/teamgooo/teamgooo-server/app/bff/notification"
	"github.com/teamgooo/teamgooo-server/app/bff/notification/notification/notificationservice"
	nsfwhelper "github.com/teamgooo/teamgooo-server/app/bff/nsfw"
	"github.com/teamgooo/teamgooo-server/app/bff/nsfw/nsfw/nsfwservice"
	passporthelper "github.com/teamgooo/teamgooo-server/app/bff/passport"
	"github.com/teamgooo/teamgooo-server/app/bff/passport/passport/passportservice"
	premiumhelper "github.com/teamgooo/teamgooo-server/app/bff/premium"
	"github.com/teamgooo/teamgooo-server/app/bff/premium/premium/premiumservice"
	privacysettingshelper "github.com/teamgooo/teamgooo-server/app/bff/privacysettings"
	"github.com/teamgooo/teamgooo-server/app/bff/privacysettings/privacysettings/privacysettingsservice"
	qrcodehelper "github.com/teamgooo/teamgooo-server/app/bff/qrcode"
	"github.com/teamgooo/teamgooo-server/app/bff/qrcode/qrcode/qrcodeservice"
	savedmessagedialogshelper "github.com/teamgooo/teamgooo-server/app/bff/savedmessagedialogs"
	"github.com/teamgooo/teamgooo-server/app/bff/savedmessagedialogs/savedmessagedialogs/savedmessagedialogsservice"
	sponsoredmessageshelper "github.com/teamgooo/teamgooo-server/app/bff/sponsoredmessages"
	"github.com/teamgooo/teamgooo-server/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"
	toshelper "github.com/teamgooo/teamgooo-server/app/bff/tos"
	"github.com/teamgooo/teamgooo-server/app/bff/tos/tos/tosservice"
	updateshelper "github.com/teamgooo/teamgooo-server/app/bff/updates"
	"github.com/teamgooo/teamgooo-server/app/bff/updates/updates/updatesservice"
	userchannelprofileshelper "github.com/teamgooo/teamgooo-server/app/bff/userchannelprofiles"
	"github.com/teamgooo/teamgooo-server/app/bff/userchannelprofiles/userchannelprofiles/userchannelprofilesservice"
	usernameshelper "github.com/teamgooo/teamgooo-server/app/bff/usernames"
	"github.com/teamgooo/teamgooo-server/app/bff/usernames/usernames/usernamesservice"
	usershelper "github.com/teamgooo/teamgooo-server/app/bff/users"
	"github.com/teamgooo/teamgooo-server/app/bff/users/users/usersservice"
	"github.com/teamgooo/teamgooo-server/pkg/net/kitex"

	"github.com/cloudwego/kitex/server"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/bff.yaml", "the config file")

type Server struct {
	kitexSrv *kitex.RpcServer
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

	s.kitexSrv = kitex.MustNewServer(
		c.RpcServerConf,
		func(s server.Server) error {
			// toshelper
			_ = tosservice.RegisterService(
				s,
				toshelper.New(toshelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// configurationhelper
			_ = configurationservice.RegisterService(
				s,
				configurationhelper.New(configurationhelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// qrcodehelper
			_ = qrcodeservice.RegisterService(
				s,
				qrcodehelper.New(
					qrcodehelper.Config{
						RpcServerConf: c.RpcServerConf,
						//KV:                c.KV,
						//UserClient:        c.BizServiceClient,
						//AuthSessionClient: c.AuthSessionClient,
						//SyncClient:        c.SyncClient,
					}))

			// miscellaneoushelper
			_ = miscellaneousservice.RegisterService(
				s,
				miscellaneoushelper.New(miscellaneoushelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// authorizationhelper
			_ = authorizationservice.RegisterService(
				s,
				authorizationhelper.New(
					authorizationhelper.Config{
						RpcServerConf: c.RpcServerConf,
						//KV:                        c.KV,
						//Code:                      c.Code,
						//UserClient:                c.BizServiceClient,
						//AuthsessionClient:         c.AuthSessionClient,
						//ChatClient:                c.BizServiceClient,
						//StatusClient:              c.StatusClient,
						//SyncClient:                c.SyncClient,
						//MsgClient:                 c.MsgClient,
						//SignInMessage:             c.SignInMessage,
						//SignInServiceNotification: c.SignInServiceNotification,
						//UsernameClient:            c.BizServiceClient,
					}))

			// premiumhelper
			_ = premiumservice.RegisterService(
				s,
				premiumhelper.New(premiumhelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// chatinviteshelper
			_ = chatinvitesservice.RegisterService(
				s,
				chatinviteshelper.New(chatinviteshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:    c.BizServiceClient,
					//ChatClient:    c.BizServiceClient,
					//MsgClient:     c.MsgClient,
					//SyncClient:    c.SyncClient,
				}))

			// chatshelper
			_ = chatsservice.RegisterService(
				s,
				chatshelper.New(chatshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:        c.BizServiceClient,
					//ChatClient:        c.BizServiceClient,
					//MsgClient:         c.MsgClient,
					//DialogClient:      c.BizServiceClient,
					//SyncClient:        c.SyncClient,
					//MediaClient:       c.MediaClient,
					//AuthsessionClient: c.AuthSessionClient,
					//IdgenClient:       c.IdgenClient,
					//MessageClient:     c.BizServiceClient,
				}))

			// fileshelper
			_ = filesservice.RegisterService(
				s,
				fileshelper.New(fileshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//DfsClient:     c.DfsClient,
					//UserClient:    c.BizServiceClient,
					//MediaClient:   c.MediaClient,
				}))

			// passporthelper
			_ = passportservice.RegisterService(
				s,
				passporthelper.New(passporthelper.Config{
					RpcServerConf: c.RpcServerConf,
					//AuthsessionClient: c.AuthSessionClient,
					//UserClient:        c.BizServiceClient,
				}))

			// updateshelper
			_ = updatesservice.RegisterService(
				s,
				updateshelper.New(updateshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UpdatesClient:     c.BizServiceClient,
					//UserClient:        c.BizServiceClient,
					//ChatClient:        c.BizServiceClient,
					//AuthsessionClient: c.AuthSessionClient,
				}))
			//
			// contactshelper
			_ = contactsservice.RegisterService(
				s,
				contactshelper.New(
					contactshelper.Config{
						RpcServerConf: c.RpcServerConf,
						//UserClient:     c.BizServiceClient,
						//ChatClient:     c.BizServiceClient,
						//UsernameClient: c.BizServiceClient,
						//SyncClient:     c.SyncClient,
					}))

			// dialogshelper
			_ = dialogsservice.RegisterService(
				s,
				dialogshelper.New(dialogshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UpdatesClient: c.BizServiceClient,
					//UserClient:    c.BizServiceClient,
					//ChatClient:    c.BizServiceClient,
					//DialogClient:  c.BizServiceClient,
					//SyncClient:    c.SyncClient,
					//MessageClient: c.BizServiceClient,
				}))

			// draftshelper
			_ = draftsservice.RegisterService(
				s,
				draftshelper.New(draftshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//DialogClient:  c.BizServiceClient,
					//UserClient:    c.BizServiceClient,
					//SyncClient:    c.SyncClient,
					//ChatClient:    c.BizServiceClient,
				}))

			// autodownloadhelper
			_ = autodownloadservice.RegisterService(
				s,
				autodownloadhelper.New(autodownloadhelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// messageshelper
			_ = messagesservice.RegisterService(
				s,
				messageshelper.New(messageshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:     c.BizServiceClient,
					//ChatClient:     c.BizServiceClient,
					//MsgClient:      c.MsgClient,
					//DialogClient:   c.BizServiceClient,
					//IdgenClient:    c.IdgenClient,
					//MessageClient:  c.BizServiceClient,
					//MediaClient:    c.MediaClient,
					//UsernameClient: c.BizServiceClient,
					//SyncClient:     c.SyncClient,
				}))

			// notificationhelper
			_ = notificationservice.RegisterService(
				s,
				notificationhelper.New(notificationhelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:    c.BizServiceClient,
					//ChatClient:    c.BizServiceClient,
					//SyncClient:    c.SyncClient,
				}))

			// usershelper
			_ = usersservice.RegisterService(
				s,
				usershelper.New(
					usershelper.Config{
						RpcServerConf: c.RpcServerConf,
						//UserClient:    c.BizServiceClient,
						//ChatClient:    c.BizServiceClient,
						//DialogClient:  c.BizServiceClient,
					}))

			// nsfwhelper
			_ = nsfwservice.RegisterService(
				s,
				nsfwhelper.New(nsfwhelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:    c.BizServiceClient,
				}))

			// sponsoredmessageshelper
			_ = sponsoredmessagesservice.RegisterService(
				s,
				sponsoredmessageshelper.New(sponsoredmessageshelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// accounthelper
			_ = accountservice.RegisterService(
				s,
				accounthelper.New(
					accounthelper.Config{
						RpcServerConf: c.RpcServerConf,
						//KV:                c.KV,
						//UserClient:        c.BizServiceClient,
						//AuthsessionClient: c.AuthSessionClient,
						//ChatClient:        c.BizServiceClient,
						//SyncClient:        c.SyncClient,
					}))

			// usernameshelper
			_ = usernamesservice.RegisterService(
				s,
				usernameshelper.New(usernameshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:     c.BizServiceClient,
					//UsernameClient: c.BizServiceClient,
					//ChatClient:     c.BizServiceClient,
					//SyncClient:     c.SyncClient,
				}))

			// privacysettingshelper
			_ = privacysettingsservice.RegisterService(
				s,
				privacysettingshelper.New(privacysettingshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UserClient:        c.BizServiceClient,
					//AuthsessionClient: c.AuthSessionClient,
					//ChatClient:        c.BizServiceClient,
					//SyncClient:        c.SyncClient,
				}))

			// savedmessagedialogshelper
			_ = savedmessagedialogsservice.RegisterService(
				s,
				savedmessagedialogshelper.New(savedmessagedialogshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//UpdatesClient: c.BizServiceClient,
					//UserClient:    c.BizServiceClient,
					//ChatClient:    c.BizServiceClient,
					//DialogClient:  c.BizServiceClient,
					//SyncClient:    c.SyncClient,
					//MessageClient: c.BizServiceClient,
				}))

			// userchannelprofilehelper
			_ = userchannelprofilesservice.RegisterService(
				s,
				userchannelprofileshelper.New(userchannelprofileshelper.Config{
					RpcServerConf: c.RpcServerConf,
					//MediaClient:   c.MediaClient,
					//UserClient:    c.BizServiceClient,
					//SyncClient:    c.SyncClient,
				}))

			return nil
		})

	// logx.Must(err)

	go func() {
		_ = s.kitexSrv.Run()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.kitexSrv.Stop()
}
