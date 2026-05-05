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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogSetChatWallpaper
// dialog.setChatWallpaper flags:# user_id:long peer_type:int peer_id:long wallpaper_id:long wallpaper_overridden:flags.0?true = Bool;
func (c *DialogCore) DialogSetChatWallpaper(in *dialog.TLDialogSetChatWallpaper) (*tg.Bool, error) {
	sourceAuth, err := c.sourcePermAuthKeyID()
	if err != nil {
		return nil, err
	}
	operationID := deterministicOperationID("set_wallpaper", in.UserId, in.PeerType, in.PeerId, in.WallpaperId, in.WallpaperOverridden)
	if err := c.svcCtx.Repo.SetPeerWallpaper(c.ctx, repository.PeerWallpaperInput{
		UserID:              in.UserId,
		PeerType:            in.PeerType,
		PeerID:              in.PeerId,
		WallpaperID:         in.WallpaperId,
		WallpaperOverridden: in.WallpaperOverridden,
		SourcePermAuthKeyID: sourceAuth,
		OperationID:         operationID,
		OutboxID:            deterministicOutboxID(operationID, "wallpaper"),
		EventType:           "dialog.wallpaperChanged",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
