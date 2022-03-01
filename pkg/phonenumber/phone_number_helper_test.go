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

package phonenumber

import (
	"fmt"
	"testing"
)

//func TestCheckPhoneNumber(t *testing.T) {
//	phone, err := CheckAndGetPhoneNumber("+639611429606")
//	fmt.Println(phone, err)
//}
//
//func TestMakePhoneNumber(t *testing.T) {
//	pNumber, err := MakePhoneNumberHelper(""+
//		"588021430", "CN")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println(pNumber.GetNormalizeDigits())
//}

func TestMakePhoneNumber2(t *testing.T) {
	// 63 969 025 1456
	// 63 995 659 1464
	pNumber, err := MakePhoneNumberHelper("+63 969 659 1464", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pNumber)
	fmt.Println(pNumber.GetNormalizeDigits(), ", ", pNumber.GetRegionCode(), ", ", pNumber.GetCountryCode())
}
