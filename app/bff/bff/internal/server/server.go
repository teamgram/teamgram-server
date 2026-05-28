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

	accounthelper "github.com/teamgram/teamgram-server/v2/app/bff/account"
	"github.com/teamgram/teamgram-server/v2/app/bff/account/account/accountservice"
	authorizationhelper "github.com/teamgram/teamgram-server/v2/app/bff/authorization"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/authorization/authorizationservice"
	autodownloadhelper "github.com/teamgram/teamgram-server/v2/app/bff/autodownload"
	"github.com/teamgram/teamgram-server/v2/app/bff/autodownload/autodownload/autodownloadservice"
	"github.com/teamgram/teamgram-server/v2/app/bff/bff/internal/config"
	chatinviteshelper "github.com/teamgram/teamgram-server/v2/app/bff/chatinvites"
	"github.com/teamgram/teamgram-server/v2/app/bff/chatinvites/chatinvites/chatinvitesservice"
	chatshelper "github.com/teamgram/teamgram-server/v2/app/bff/chats"
	"github.com/teamgram/teamgram-server/v2/app/bff/chats/chats/chatsservice"
	configurationhelper "github.com/teamgram/teamgram-server/v2/app/bff/configuration"
	"github.com/teamgram/teamgram-server/v2/app/bff/configuration/configuration/configurationservice"
	contactshelper "github.com/teamgram/teamgram-server/v2/app/bff/contacts"
	"github.com/teamgram/teamgram-server/v2/app/bff/contacts/contacts/contactsservice"
	dialogshelper "github.com/teamgram/teamgram-server/v2/app/bff/dialogs"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/dialogs/dialogsservice"
	draftshelper "github.com/teamgram/teamgram-server/v2/app/bff/drafts"
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/drafts/draftsservice"
	fileshelper "github.com/teamgram/teamgram-server/v2/app/bff/files"
	"github.com/teamgram/teamgram-server/v2/app/bff/files/files/filesservice"
	messageshelper "github.com/teamgram/teamgram-server/v2/app/bff/messages"
	"github.com/teamgram/teamgram-server/v2/app/bff/messages/messages/messagesservice"
	miscellaneoushelper "github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous"
	"github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
	notificationhelper "github.com/teamgram/teamgram-server/v2/app/bff/notification"
	"github.com/teamgram/teamgram-server/v2/app/bff/notification/notification/notificationservice"
	nsfwhelper "github.com/teamgram/teamgram-server/v2/app/bff/nsfw"
	"github.com/teamgram/teamgram-server/v2/app/bff/nsfw/nsfw/nsfwservice"
	passporthelper "github.com/teamgram/teamgram-server/v2/app/bff/passport"
	"github.com/teamgram/teamgram-server/v2/app/bff/passport/passport/passportservice"
	premiumhelper "github.com/teamgram/teamgram-server/v2/app/bff/premium"
	"github.com/teamgram/teamgram-server/v2/app/bff/premium/premium/premiumservice"
	privacysettingshelper "github.com/teamgram/teamgram-server/v2/app/bff/privacysettings"
	"github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/privacysettings/privacysettingsservice"
	qrcodehelper "github.com/teamgram/teamgram-server/v2/app/bff/qrcode"
	"github.com/teamgram/teamgram-server/v2/app/bff/qrcode/qrcode/qrcodeservice"
	savedmessagedialogshelper "github.com/teamgram/teamgram-server/v2/app/bff/savedmessagedialogs"
	"github.com/teamgram/teamgram-server/v2/app/bff/savedmessagedialogs/savedmessagedialogs/savedmessagedialogsservice"
	sponsoredmessageshelper "github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages"
	"github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"
	toshelper "github.com/teamgram/teamgram-server/v2/app/bff/tos"
	"github.com/teamgram/teamgram-server/v2/app/bff/tos/tos/tosservice"
	updateshelper "github.com/teamgram/teamgram-server/v2/app/bff/updates"
	"github.com/teamgram/teamgram-server/v2/app/bff/updates/updates/updatesservice"
	userchannelprofileshelper "github.com/teamgram/teamgram-server/v2/app/bff/userchannelprofiles"
	"github.com/teamgram/teamgram-server/v2/app/bff/userchannelprofiles/userchannelprofiles/userchannelprofilesservice"
	usernameshelper "github.com/teamgram/teamgram-server/v2/app/bff/usernames"
	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/usernames/usernamesservice"
	usershelper "github.com/teamgram/teamgram-server/v2/app/bff/users"
	"github.com/teamgram/teamgram-server/v2/app/bff/users/users/usersservice"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/cloudwego/kitex/server"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/bff.yaml", "the config file")

type Server struct {
	kitexSrv *kitex.RpcServer
}

func withServiceName(c kitex.RpcClientConf, serviceName string) kitex.RpcClientConf {
	c.ServiceName = serviceName
	return c
}

func buildChatInvitesConfig(c config.Config) chatinviteshelper.Config {
	return chatinviteshelper.Config{
		RpcServerConf: c.RpcServerConf,
		ChatClient:    withServiceName(c.BizServiceClient, "RPCChat"),
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildChatsConfig(c config.Config) chatshelper.Config {
	return chatshelper.Config{
		RpcServerConf: c.RpcServerConf,
		ChatClient:    withServiceName(c.BizServiceClient, "RPCChat"),
		DialogClient:  withServiceName(c.BizServiceClient, "RPCDialog"),
		MsgClient:     withServiceName(c.MsgClient, "RPCMsg"),
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildContactsConfig(c config.Config) contactshelper.Config {
	return contactshelper.Config{
		RpcServerConf: c.RpcServerConf,
		ChatClient:    withServiceName(c.BizServiceClient, "RPCChat"),
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildDialogsConfig(c config.Config) dialogshelper.Config {
	return dialogshelper.Config{
		RpcServerConf:            c.RpcServerConf,
		ChatClient:               withServiceName(c.BizServiceClient, "RPCChat"),
		DialogClient:             withServiceName(c.BizServiceClient, "RPCDialog"),
		MsgClient:                withServiceName(c.MsgClient, "RPCMsg"),
		SyncClient:               withServiceName(c.SyncClient, "RPCSync"),
		UserupdatesClient:        withServiceName(c.UserupdatesClient, "RPCUserupdates"),
		UserClient:               withServiceName(c.BizServiceClient, "RPCUser"),
		TypingMinIntervalSeconds: c.TypingMinIntervalSeconds,
	}
}

func buildDraftsConfig(c config.Config) draftshelper.Config {
	return draftshelper.Config{
		RpcServerConf: c.RpcServerConf,
		DialogClient:  withServiceName(c.BizServiceClient, "RPCDialog"),
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
		ChatClient:    withServiceName(c.BizServiceClient, "RPCChat"),
	}
}

func buildMessagesConfig(c config.Config) messageshelper.Config {
	return messageshelper.Config{
		RpcServerConf:     c.RpcServerConf,
		ChatClient:        withServiceName(c.BizServiceClient, "RPCChat"),
		IdgenClient:       withServiceName(c.IdgenClient, "RPCIdgen"),
		MsgClient:         withServiceName(c.MsgClient, "RPCMsg"),
		MediaClient:       withServiceName(c.MediaClient, "RPCMedia"),
		UserupdatesClient: withServiceName(c.UserupdatesClient, "RPCUserupdates"),
		UserClient:        withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildAuthorizationConfig(c config.Config) authorizationhelper.Config {
	return authorizationhelper.Config{
		RpcServerConf:     c.RpcServerConf,
		AuthsessionClient: withServiceName(c.AuthSessionClient, "RPCAuthsession"),
		UserClient:        withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildUpdatesConfig(c config.Config) updateshelper.Config {
	return updateshelper.Config{
		RpcServerConf:     c.RpcServerConf,
		UserupdatesClient: withServiceName(c.UserupdatesClient, "RPCUserupdates"),
		UserClient:        withServiceName(c.BizServiceClient, "RPCUser"),
		ChatClient:        withServiceName(c.BizServiceClient, "RPCChat"),
	}
}

func buildUsersConfig(c config.Config) usershelper.Config {
	return usershelper.Config{
		RpcServerConf: c.RpcServerConf,
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildNsfwConfig(c config.Config) nsfwhelper.Config {
	return nsfwhelper.Config{
		RpcServerConf: c.RpcServerConf,
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
	}
}

func buildAccountConfig(c config.Config) accounthelper.Config {
	return accounthelper.Config{
		RpcServerConf:     c.RpcServerConf,
		UserClient:        withServiceName(c.BizServiceClient, "RPCUser"),
		AuthsessionClient: withServiceName(c.AuthSessionClient, "RPCAuthsession"),
		CodeClient:        withServiceName(c.BizServiceClient, "RPCCode"),
	}
}

func buildUsernamesConfig(c config.Config) usernameshelper.Config {
	return usernameshelper.Config{
		RpcServerConf: c.RpcServerConf,
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
		ChatClient:    withServiceName(c.BizServiceClient, "RPCChat"),
	}
}

func buildPrivacySettingsConfig(c config.Config) privacysettingshelper.Config {
	return privacysettingshelper.Config{
		RpcServerConf:     c.RpcServerConf,
		UserClient:        withServiceName(c.BizServiceClient, "RPCUser"),
		ChatClient:        withServiceName(c.BizServiceClient, "RPCChat"),
		AuthsessionClient: withServiceName(c.AuthSessionClient, "RPCAuthsession"),
	}
}

func buildSavedMessageDialogsConfig(c config.Config) savedmessagedialogshelper.Config {
	return savedmessagedialogshelper.Config{
		RpcServerConf: c.RpcServerConf,
		DialogClient:  withServiceName(c.BizServiceClient, "RPCDialog"),
	}
}

func buildUserChannelProfilesConfig(c config.Config) userchannelprofileshelper.Config {
	return userchannelprofileshelper.Config{
		RpcServerConf: c.RpcServerConf,
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
		MediaClient:   withServiceName(c.MediaClient, "RPCMedia"),
	}
}

func buildFilesConfig(c config.Config) fileshelper.Config {
	return fileshelper.Config{
		RpcServerConf: c.RpcServerConf,
		DfsClient:     withServiceName(c.DfsClient, "RPCDfs"),
		MediaClient:   withServiceName(c.MediaClient, "RPCMedia"),
		UserClient:    withServiceName(c.BizServiceClient, "RPCUser"),
	}
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
				authorizationhelper.New(buildAuthorizationConfig(c)))

			// premiumhelper
			_ = premiumservice.RegisterService(
				s,
				premiumhelper.New(premiumhelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// chatinviteshelper
			_ = chatinvitesservice.RegisterService(
				s,
				chatinviteshelper.New(buildChatInvitesConfig(c)))

			// chatshelper
			_ = chatsservice.RegisterService(
				s,
				chatshelper.New(buildChatsConfig(c)))

			// fileshelper
			_ = filesservice.RegisterService(
				s,
				fileshelper.New(buildFilesConfig(c)))

			// passporthelper
			_ = passportservice.RegisterService(
				s,
				passporthelper.New(passporthelper.Config{
					RpcServerConf:     c.RpcServerConf,
					AuthsessionClient: c.AuthSessionClient,
					UserClient:        c.BizServiceClient,
				}))

			// updateshelper
			_ = updatesservice.RegisterService(
				s,
				updateshelper.New(buildUpdatesConfig(c)))
			//
			// contactshelper
			_ = contactsservice.RegisterService(
				s,
				contactshelper.New(buildContactsConfig(c)))

			// dialogshelper
			_ = dialogsservice.RegisterService(
				s,
				dialogshelper.New(buildDialogsConfig(c)))

			// draftshelper
			_ = draftsservice.RegisterService(
				s,
				draftshelper.New(buildDraftsConfig(c)))

			// autodownloadhelper
			_ = autodownloadservice.RegisterService(
				s,
				autodownloadhelper.New(autodownloadhelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// messageshelper
			_ = messagesservice.RegisterService(
				s,
				messageshelper.New(buildMessagesConfig(c)))

			// notificationhelper
			_ = notificationservice.RegisterService(
				s,
				notificationhelper.New(notificationhelper.Config{
					RpcServerConf: c.RpcServerConf,
					UserClient:    c.BizServiceClient,
					ChatClient:    c.BizServiceClient,
				}, nil))

			// usershelper
			_ = usersservice.RegisterService(
				s,
				usershelper.New(buildUsersConfig(c)))

			// nsfwhelper
			_ = nsfwservice.RegisterService(
				s,
				nsfwhelper.New(buildNsfwConfig(c)))

			// sponsoredmessageshelper
			_ = sponsoredmessagesservice.RegisterService(
				s,
				sponsoredmessageshelper.New(sponsoredmessageshelper.Config{
					RpcServerConf: c.RpcServerConf,
				}))

			// accounthelper
			_ = accountservice.RegisterService(
				s,
				accounthelper.New(buildAccountConfig(c), nil))

			// usernameshelper
			_ = usernamesservice.RegisterService(
				s,
				usernameshelper.New(buildUsernamesConfig(c)))

			// privacysettingshelper
			_ = privacysettingsservice.RegisterService(
				s,
				privacysettingshelper.New(buildPrivacySettingsConfig(c)))

			// savedmessagedialogshelper
			_ = savedmessagedialogsservice.RegisterService(
				s,
				savedmessagedialogshelper.New(buildSavedMessageDialogsConfig(c)))

			// userchannelprofilehelper
			_ = userchannelprofilesservice.RegisterService(
				s,
				userchannelprofileshelper.New(buildUserChannelProfilesConfig(c)))

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
