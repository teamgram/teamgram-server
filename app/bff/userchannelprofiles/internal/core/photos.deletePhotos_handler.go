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
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// PhotosDeletePhotos
// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (c *UserChannelProfilesCore) PhotosDeletePhotos(in *tg.TLPhotosDeletePhotos) (*tg.VectorLong, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	deleteIDs := make([]int64, 0, len(in.Id))
	for _, inputPhoto := range in.Id {
		id, err := photoIDFromInputPhoto(inputPhoto)
		if err != nil {
			return nil, err
		}
		deleteIDs = append(deleteIDs, id)
	}

	mainID, err := c.svcCtx.Repo.UserClient.UserDeleteProfilePhotos(c.ctx, &userpb.TLUserDeleteProfilePhotos{
		UserId: selfID,
		Id:     deleteIDs,
	})
	if err != nil {
		return nil, err
	}
	if mainID != nil && mainID.V > 0 {
		if err := requireMediaClient(c); err != nil {
			return nil, err
		}
		if _, err = c.svcCtx.Repo.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
			PhotoId: mainID.V,
		}); err != nil {
			return nil, err
		}
	}
	// TODO(v2 userchannelprofiles): sync delivery is intentionally not migrated here; route profile photo updates through userupdates/gateway when the V2 delivery contract is defined.

	return &tg.VectorLong{Datas: deleteIDs}, nil
}
