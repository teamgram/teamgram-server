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

package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

// UserGetUserProjectionBundle
// user.getUserProjectionBundle flags:# with_facts:flags.0?true viewer_user_ids:Vector<long> target_user_ids:Vector<long> = UserProjectionBundle;
func (c *UserCore) UserGetUserProjectionBundle(in *user.TLUserGetUserProjectionBundle) (*user.UserProjectionBundle, error) {
	if in == nil || len(in.ViewerUserIds) == 0 {
		return nil, user.ErrUserInvalidArgument
	}
	repo := c.svcCtx.UserProjectionRepo
	if repo == nil {
		repo = c.svcCtx.Repo
	}
	if repo == nil {
		return nil, user.ErrUserStorage
	}
	bundle, err := repo.GetUserProjectionBundle(c.ctx, in.ViewerUserIds, in.TargetUserIds, in.WithFacts)
	if err != nil {
		return nil, err
	}
	return user.MakeTLUserProjectionBundle(&user.TLUserProjectionBundle{
		Facts:          bundle.Facts,
		ViewerUsers:    repositoryViewerUsersToTL(bundle.ViewerUsers),
		MissingUserIds: bundle.MissingUserIds,
	}).ToUserProjectionBundle(), nil
}

func repositoryViewerUsersToTL(in []repository.ViewerUsers) []user.ViewerUsersClazz {
	out := make([]user.ViewerUsersClazz, 0, len(in))
	for _, item := range in {
		out = append(out, user.MakeTLViewerUsers(&user.TLViewerUsers{
			ViewerUserId: item.ViewerUserId,
			Users:        item.Users,
		}).ToViewerUsers())
	}
	return out
}
