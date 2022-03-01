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

// type UserType int

const (
	UserTypeUnknown = 0
	UserTypeDeleted = 1
	UserTypeRegular = 2
	UserTypeBot     = 3
	UserTypeService = 4
	UserTypeTest    = 5
)

//func IsSupportId(id int32) bool {
//	return id/1000 == 777 || id == 333000 ||
//		id == 4240000 || id == 4240000 || id == 4244000 ||
//		id == 4245000 || id == 4246000 || id == 410000 ||
//		id == 420000 || id == 431000 || id == 431415000 ||
//		id == 434000 || id == 4243000 || id == 439000 ||
//		id == 449000 || id == 450000 || id == 452000 ||
//		id == 454000 || id == 4254000 || id == 455000 ||
//		id == 460000 || id == 470000 || id == 479000 ||
//		id == 796000 || id == 482000 || id == 490000 ||
//		id == 496000 || id == 497000 || id == 498000 ||
//		id == 4298000
//}
//
//func IsUserTest(id int32) bool {
//	return id >= 200 && id < 500
//}
