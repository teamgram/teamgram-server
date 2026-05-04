// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package svc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	_ "github.com/teamgram/teamgram-server/v2/app/bff/account/account/accountservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/authorization/authorization/authorizationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/autodownload/autodownload/autodownloadservice"
	bffproxyclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/chatinvites/chatinvites/chatinvitesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/chats/chats/chatsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/configuration/configuration/configurationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/contacts/contacts/contactsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/dialogs/dialogs/dialogsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/drafts/drafts/draftsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/files/files/filesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/messages/messages/messagesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/notification/notification/notificationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/nsfw/nsfw/nsfwservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/passport/passport/passportservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/premium/premium/premiumservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/privacysettings/privacysettingsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/qrcode/qrcode/qrcodeservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/savedmessagedialogs/savedmessagedialogs/savedmessagedialogsservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/sponsoredmessages/sponsoredmessages/sponsoredmessagesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/tos/tos/tosservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/updates/updates/updatesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/userchannelprofiles/userchannelprofiles/userchannelprofilesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/usernames/usernames/usernamesservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/users/users/usersservice"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/config"
	gatewaypresence "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/presence"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/push"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/repository"
	presenceclient "github.com/teamgram/teamgram-server/v2/app/service/presence/client"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type ServiceContext struct {
	Config            config.Config
	GatewayGeneration string
	Repo              *repository.Repository
	BFF               *bffproxyclient.BFFProxyClient2
	Push              *push.LocalWriter
	Presence          *gatewaypresence.Registrar
}

func NewServiceContext(c config.Config) *ServiceContext {
	generation := newGatewayGeneration()
	var presenceRegistrar *gatewaypresence.Registrar
	if presenceClientConfigured(c.PresenceClient) {
		presenceRegistrar = gatewaypresence.NewRegistrar(gatewaypresence.Config{
			GatewayID:                  c.GatewayId,
			GatewayGeneration:          generation,
			GatewayRPCAddr:             c.AdvertiseRpcAddr,
			RefreshMinInterval:         time.Duration(c.PresenceRefreshMinIntervalSeconds) * time.Second,
			RefreshRetryMinInterval:    time.Duration(c.PresenceRefreshRetryMinIntervalSeconds) * time.Second,
			RefreshScanInterval:        time.Duration(c.PresenceRefreshScanIntervalSeconds) * time.Second,
			ShutdownOfflineDeadline:    time.Duration(c.GatewayShutdownPresenceOfflineDeadlineSeconds) * time.Second,
			ShutdownOfflineMaxSessions: c.GatewayShutdownPresenceOfflineMaxSessions,
		}, generatedPresenceClient{client: presenceclient.NewPresenceClient(presenceclient.MustNewKitexClient(c.PresenceClient))}, time.Now)
	}
	return &ServiceContext{
		Config:            c,
		GatewayGeneration: generation,
		Repo:              repository.NewRepository(c),
		BFF:               bffproxyclient.NewBFFProxyClient2(c.BffClient.Clients),
		Push:              push.NewLocalWriter(),
		Presence:          presenceRegistrar,
	}
}
func (s *ServiceContext) Close() error {
	if s == nil || s.Repo == nil {
		return nil
	}
	return s.Repo.Close()
}

type generatedPresenceClient struct {
	client presenceclient.PresenceClient
}

func (c generatedPresenceClient) SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession) error {
	_, err := c.client.PresenceSetSessionOnline(ctx, &presencepb.TLPresenceSetSessionOnline{Session: session})
	return err
}

func (c generatedPresenceClient) SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error {
	_, err := c.client.PresenceSetSessionOffline(ctx, &presencepb.TLPresenceSetSessionOffline{
		UserId:    userID,
		AuthKeyId: authKeyID,
		SessionId: sessionID,
	})
	return err
}

func newGatewayGeneration() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err == nil {
		return hex.EncodeToString(b[:])
	}
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func presenceClientConfigured(c kitex.RpcClientConf) bool {
	return c.DestService != "" || len(c.Endpoints) > 0 || c.HasEtcd() || c.Target != ""
}
