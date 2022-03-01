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

package user

import (
	"github.com/teamgram/proto/mtproto"
)

type BlockedList []*mtproto.PeerBlocked

// Len len
func (m BlockedList) Len() int {
	return len(m)
}
func (m BlockedList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m BlockedList) Less(i, j int) bool {
	// TODO(@benqi): if date[i] == date[j]
	return m[i].Date < m[j].Date
}

type Contact struct {
	SelfUserId    int32  `json:"self_user_id"`
	ContactUserId int32  `json:"contact_user_id"`
	PhoneNumber   string `json:"phone_number"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MutualContact bool   `json:"mutual_contact"`
}
