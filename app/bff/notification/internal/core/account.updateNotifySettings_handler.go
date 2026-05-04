// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/bff/notification/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountUpdateNotifySettings
// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (c *NotificationCore) AccountUpdateNotifySettings(in *tg.TLAccountUpdateNotifySettings) (*tg.Bool, error) {
	peerUtil := fromInputNotifyPeer(c.MD.UserId, in.Peer)

	// Validate peer existence for specific peer types
	switch peerUtil.PeerType {
	case tg.PEER_CHAT:
		_, err := c.svcCtx.Repo.ChatClient.ChatGetMutableChat(c.ctx, &repository.GetMutableChat{
			ChatId: peerUtil.PeerId,
		})
		if err != nil {
			c.Logger.Errorf("account.updateNotifySettings - error: chat %d not found: %v", peerUtil.PeerId, err)
			return nil, err
		}
	case tg.PEER_CHANNEL:
		if c.svcCtx.Plugin != nil {
			_, err := c.svcCtx.Plugin.GetChannelById(c.ctx, c.MD.UserId, peerUtil.PeerId)
			if err != nil {
				c.Logger.Errorf("account.updateNotifySettings - error: channel %d not found: %v", peerUtil.PeerId, err)
				return nil, err
			}
		}
	}

	// Convert InputPeerNotifySettings to PeerNotifySettings
	// InputPeerNotifySettingsClazz is *TLInputPeerNotifySettings, so we use it directly.
	inputSettings := in.Settings
	if inputSettings == nil {
		inputSettings = &tg.TLInputPeerNotifySettings{}
	}

	peerSettings := tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{
		ShowPreviews:        inputSettings.ShowPreviews,
		Silent:              inputSettings.Silent,
		MuteUntil:           inputSettings.MuteUntil,
		IosSound:            inputSettings.Sound,
		AndroidSound:        inputSettings.Sound,
		OtherSound:          inputSettings.Sound,
		StoriesMuted:        inputSettings.StoriesMuted,
		StoriesHideSender:   inputSettings.StoriesHideSender,
		StoriesIosSound:     inputSettings.StoriesSound,
		StoriesAndroidSound: inputSettings.StoriesSound,
		StoriesOtherSound:   inputSettings.StoriesSound,
	})

	_, err := c.svcCtx.Repo.UserClient.UserSetNotifySettings(c.ctx, &repository.SetNotifySettings{
		UserId:   c.MD.UserId,
		PeerType: peerUtil.PeerType,
		PeerId:   peerUtil.PeerId,
		Settings: peerSettings,
	})
	if err != nil {
		c.Logger.Errorf("account.updateNotifySettings - error: %v", err)
		return nil, err
	}

	// TODO: sync.SyncUpdatesNotMe

	return tg.BoolTrue, nil
}
