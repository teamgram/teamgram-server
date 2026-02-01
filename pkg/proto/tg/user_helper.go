// Copyright 2024 Teamgram Authors
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

package tg

func ToUserIdByInput(userSelfId int64, inputUser *InputUser) (id int64) {
	inputUser.Match(
		func(c *TLInputUserEmpty) interface{} {
			id = 0

			return nil
		},
		func(c *TLInputUserSelf) interface{} {
			id = userSelfId

			return nil
		},
		func(c *TLInputUser) interface{} {
			id = c.UserId

			return nil
		})

	return
}

func ToUserIdListByInput(userSelfId int64, inputUsers []*InputUser) []int64 {
	idList := make([]int64, 0, len(inputUsers))
	for _, user := range inputUsers {
		id := ToUserIdByInput(userSelfId, user)
		if id > 0 {
			idList = append(idList, id)
		} else {
			// ignore in
		}
	}
	return idList
}

func isUserDeleted(user *User) bool {
	rV := false

	if user != nil {
		user.Match(
			func(c *TLUserEmpty) interface{} {
				rV = true

				return nil
			},
			func(c *TLUser) interface{} {
				rV = c.Deleted

				return nil
			})
	} else {
		rV = true
	}

	return rV
}

func isUserContact(user *User) bool {
	if user2, ok := user.ToUser(); ok {
		return user2.Contact || user2.MutualContact
	} else {
		return false
	}
}

func isUserSelf(user *User) bool {
	if user2, ok := user.ToUser(); ok {
		return user2.Self
	} else {
		return false
	}
}

func GetUserName(user *User) (name string) {
	if user == nil {
		name = "Deleted Account"
	} else {
		user.Match(
			func(c *TLUserEmpty) interface{} {
				name = "Deleted Account"

				return nil
			},
			func(c *TLUser) interface{} {
				firstName := GetFlagsString(c.FirstName)
				lastName := GetFlagsString(c.LastName)

				if firstName == "" && lastName == "" {
					name = ""
				} else if firstName == "" {
					name = lastName
				} else if lastName == "" {
					name = firstName
				} else {
					name = firstName + " " + lastName
				}

				return nil
			})
	}

	return
}
