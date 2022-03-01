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

package main

import (
	"fmt"

	"github.com/teamgram/teamgram-server/pkg/mention"
)

var testRuneStr = []rune(`1) a b (@ababab)
Телефон: 8613745848668 | Код: 12345
Управление: /u_8613745848668 

2) BanMePlease  (@banmeplease)
Телефон: 8613486832450 | Код: 11111
Управление: /u_8613486832450 

3) Daniil Sokol (@dsokol) [✅]
Телефон: 8613692174272 | Код: 14441
Управление: /u_8613692174272 

4) Ivan Ivanov (@IvanIvanov)
Телефон: 79997654321 | Код: 97654
Управление: /u_79997654321 

5) Without Code (@NoCodeTestUser)
Телефон: 8613019689412 | Код: 27483
Управление: /u_8613019689412 

6) Petya Sidorov (@psidorov)
Телефон: 8613000000001 | Код: 11333
Управление: /u_8613000000001 

7) Test User2 (@TestUser0002)
Телефон: 79991234568 | Код: 97531
Управление: /u_79991234568 

😍 TestUser3  (@testuser0003)
Телефон: 8613170211337 | Код: 62034
Управление: /u_8613170211337 

9) TestUser4  (@testuser0004) [✅]
Телефон: 8613678122163 | Код: 84387
Управление: /u_8613678122163 

10) TestUser1  (@TestUser1)
Телефон: 79991234567 | Код: 13579
Управление: /u_79991234567`)

var (
	utf8TestStr = `1) a b (@ababab)
Телефон: 8613745848668 | Код: 12345
Управление: /u_8613745848668 

2) BanMePlease  (@banmeplease)
Телефон: 8613486832450 | Код: 12689
Управление: /u_8613486832450 

3) Daniil Sokol (@dsokol) [✅]
Телефон: 8613692174272 | Код: 14441
Управление: /u_8613692174272 

4) l l (@fthftjul)
Телефон: 8613181543558 | Код: 32468
Управление: /u_8613181543558 

5) Ivan Ivanov (@IvanIvanov)
Телефон: 79997654321 | Код: 97654
Управление: /u_79997654321 

6) Without Code (@NoCodeTestUser)
Телефон: 8613019689412 | Код: 27483
Управление: /u_8613019689412 

7) Petya Sidorov (@psidorov)
Телефон: 8613000000001 | Код: 11333
Управление: /u_8613000000001 

😍 Test User2 (@TestUser0002)
Телефон: 79991234568 | Код: 97531
Управление: /u_79991234568 

9) TestUser4  (@testuser0004)
Телефон: 8613678122163 | Код: 84387
Управление: /u_8613678122163 

10) TestUser1  (@TestUser1)
Телефон: 79991234567 | Код: 13579
Управление: /u_79991234567 

11) TestUser3  (@Usertest333)
Телефон: 8613170211337 | Код: 66666
Управление: /u_8613170211337

@TestUser0002`

	utf8TestStr2 = `1) a b (@ababab)
Телефон: 8613745848668 | Код: 12345
Управление: /u_8613745848668 

2) BanMePlease  (@banmeplease)
Телефон: 8613486832450 | Код: 12689
Управление: /u_8613486832450 

3) Daniil Sokol (@dsokol) [✅]
Телефон: 8613692174272 | Код: 14441
Управление: /u_8613692174272 

4) l l (@fthftjul)
Телефон: 8613181543558 | Код: 32468
Управление: /u_8613181543558 

5) Ivan Ivanov (@IvanIvanov)
Телефон: 79997654321 | Код: 97654
Управление: /u_79997654321 

6) Without Code (@NoCodeTestUser)
Телефон: 8613019689412 | Код: 27483
Управление: /u_8613019689412 

7) Petya Sidorov (@psidorov)
Телефон: 8613000000001 | Код: 11333
Управление: /u_8613000000001 

😍 Test User2 (@TestUser0002)
Телефон: 79991234568 | Код: 97531
Управление: /u_79991234568 

9) TestUser4  (@testuser0004)
Телефон: 8613678122163 | Код: 84387
Управление: /u_8613678122163 

10) TestUser1  (@TestUser1)
Телефон: 79991234567 | Код: 13579
Управление: /u_79991234567 

11) TestUser3  (@Usertest333)
Телефон: 8613170211337 | Код: 66666
Управление: /u_8613170211337`
)

func main() {
	//utfEncodedString := utf16.Encode(testRuneStr)
	//
	//tags := mention.GetTagsAsUniqueUTF16Strings('@', utfEncodedString, '(', ')')
	//for _, tag := range tags {
	//	fmt.Println(string(utf16.Decode(tag)))
	//}
	//tags = mention.GetTagsAsUniqueUTF16Strings('/', utfEncodedString)
	//for _, tag := range tags {
	//	fmt.Println(string(utf16.Decode(tag)))
	//}
	//
	//tags2 := mention.GetUTF16Tags('@', utfEncodedString, '(', ')')
	//for _, tag := range tags2 {
	//	fmt.Println(tag)
	//}
	//tags2 = mention.GetUTF16Tags('/', utfEncodedString)
	//for _, tag := range tags {
	//	fmt.Println(tag)
	//}

	fmt.Println(len(utf8TestStr))
	utf16TestStr := mention.EncodeStringToUTF16(utf8TestStr)
	fmt.Println(len(utf16TestStr))
	_ = utf16TestStr
	idxList := mention.EncodeStringToUTF16Index(utf8TestStr)
	_ = idxList
	fmt.Println(idxList)
	var tags []mention.Tag

	//tags = mention.GetTags('@', utf8TestStr2, '(', ')')
	//for _, tag := range tags {
	//	//fmt.Println(tag.Index)
	//	//o0 := sort.SearchInts(idxList, tag.Index)
	//	//o1 := sort.SearchInts(idxList, tag.Index+len(tag.Tag)+1)
	//	//fmt.Println(o0, " ==> ", o1)
	//
	//	fmt.Println(tag.Index)
	//	fmt.Println(tag.Index + len(tag.Tag) + 1)
	//	fmt.Println(idxList[tag.Index])
	//	fmt.Println(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index])
	//
	//	fmt.Println(tag.Tag)
	//	fmt.Println(tag.Index)
	//	fmt.Println(mention.DecodeUTF16ToString(utf16TestStr[idxList[tag.Index] : idxList[tag.Index]+len(tag.Tag)+1]))
	//}

	tags = mention.GetTags('/', utf8TestStr2)
	for _, tag := range tags {
		fmt.Println(tag.Index)
		//o0 := sort.SearchInts(idxList, tag.Index)
		//o1 := sort.SearchInts(idxList, tag.Index+len(tag.Tag)+1)
		//fmt.Println(o0, " ==> ", o1)

		fmt.Println(tag.Index)
		fmt.Println(tag.Index + len(tag.Tag) + 1)
		fmt.Println(idxList[tag.Index])
		fmt.Println(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index])

		fmt.Println(tag.Tag)
		fmt.Println(mention.DecodeUTF16ToString(utf16TestStr[idxList[tag.Index] : idxList[tag.Index]+len(tag.Tag)+1]))
	}

}
