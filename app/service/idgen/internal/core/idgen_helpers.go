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
	"fmt"

	idgenpb "github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
)

// maxNextIdsNum bounds idgen.nextIds batch size to keep RPC payloads
// predictable. Repository.NextIDs itself does not enforce this; the cap
// belongs to the service-level use-case policy enforced by core.
const maxNextIdsNum = 100

// validateNextIdsNum rejects num values that are negative or exceed the
// per-call cap. The check is mirrored for both idgen.nextIds and the
// inputIds branch of idgen.getNextIdValList.
func validateNextIdsNum(num int32) error {
	if num < 0 || num > maxNextIdsNum {
		return fmt.Errorf("%w: next ids num %d out of range [0,%d]", idgenpb.ErrInvalidArgument, num, maxNextIdsNum)
	}
	return nil
}

// nextIDs validates num against the service-level cap and delegates to
// Repository for id generation.
func (c *IdgenCore) nextIDs(num int32) ([]int64, error) {
	if err := validateNextIdsNum(num); err != nil {
		return nil, err
	}
	return c.svcCtx.Repo.NextIDs(num), nil
}
