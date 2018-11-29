// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package server

import (
	"fmt"
	"testing"
)

func TestClientSessionManager(t *testing.T) {
	s := newClientSessionManager(100000)
	s.Start()

	fmt.Println("ready.")
	for i := 0; i < 10; i++ {
		s.onSessionData(&sessionData{ClientConnID{1, 1, 1}, nil, []byte{1}})
	}

	s.Stop()
}


func CheckClassID(classId int32) bool {
	var registers2 = map[int32]int32{}
	_, ok := registers2[classId]
	return ok
}
