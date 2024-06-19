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

package plugin

import (
	"context"

	"github.com/teamgram/proto/mtproto"
)

type StoryPlugin interface {
	// GetStoriesPinnedAvailable
	// stories_pinned_available	flags.26?true	Whether this user has some pinned stories.
	GetStoriesPinnedAvailable(ctx context.Context, peerUserId, toSelfUserId int64) bool
	// GetBlockedMyStoriesFrom
	// blocked_my_stories_from	flags.27?true	Whether we've blocked this user, preventing them from seeing our stories ».
	GetBlockedMyStoriesFrom(ctx context.Context, peerUserId, toSelfUserId int64) bool
	// GetActiveStories
	// stories	flags.25?PeerStories	Active stories »
	GetActiveStories(ctx context.Context, peerUserId, toSelfUserId int64) *mtproto.PeerStories
}
