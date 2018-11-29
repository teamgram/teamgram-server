// Copyright (C) 2012-2013 Fuchun(Aying) All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//           http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package random2

import (
	"fmt"
	"regexp"
	"testing"
)

func assertTrue(t *testing.T, expr bool, arg1 string, arg2 string) {
	if expr {
		t.Log(arg1)
	} else {
		t.Error(arg2)
	}
}

func assertEquals(t *testing.T, arg1, arg2 interface{}, falseMsg string, params ...interface{}) {
	if arg1 == arg2 {
		t.Log("PASS")
	} else {
		if len(params) == 0 {
			t.Error(falseMsg)
		} else {
			t.Errorf(falseMsg, params)
		}
	}
}

func TestRandomString(t *testing.T) {

	// random utf8 string
	str1 := RandomString(21)
	assertEquals(t, len([]rune(str1)), 21, "RandomString(21) length")

	str2 := RandomString(21)
	assertEquals(t, len([]rune(str2)), 21, "RandomString(21) length")
	assertTrue(t, func() bool {
		str1Rune := []rune(str1)
		str2Rune := []rune(str2)
		for i := 0; i < len(str1Rune); i++ {
			if str1Rune[i] != str2Rune[i] {
				return false
			}
		}
		return true
	}(), "PASS", "str1 != str2")

	// random ascii
	str1 = RandomAscii(21)
	assertEquals(t, len(str1), 21, "RandomAscii(21) length")
	for _, r := range []rune(str1) {
		assertTrue(t, r >= 32 && r <= 127, "PASS", "char between 32 and 127")
	}
	str2 = RandomAscii(21)
	assertTrue(t, str1 != str2, "PASS", "str1 != str2")
	fmt.Printf("RandomAscii(21):\nstr1=%s\nstr2=%s\n", str1, str2)

	// random alphabetic
	str1 = RandomAlphabetic(21)
	assertEquals(t, len(str1), 21, "RandomAlphabetic(21) length")
	assertTrue(t, func() bool {
		str1Rune := []rune(str1)
		all_alphabetic := true
		for i := 0; i < len(str1Rune); i++ {
			ch := str1Rune[i]
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				continue
			} else {
				all_alphabetic = false
				break
			}
		}
		return all_alphabetic
	}(), "PASS", "str1 is all alphabetic")
	str2 = RandomAlphabetic(21)
	assertTrue(t, str1 != str2, "PASS", "str1 != str2")
	fmt.Printf("RandomAlphabetic(21):\nstr1=%s\nstr2=%s\n", str1, str2)

	// random numeric
	str1 = RandomNumeric(21)
	assertEquals(t, len(str1), 21, "RandomNumeric(21) length")
	assertTrue(t, func() bool {
		numeric, _ := regexp.MatchString("^\\d+$", str1)
		return numeric
	}(), "PASS", "str1 is all numeric")
	str2 = RandomNumeric(21)
	assertTrue(t, str1 != str2, "PASS", "str1 != str2")
	fmt.Printf("RandomNumeric(21):\nstr1=%s\nstr2=%s\n", str1, str2)

	// random alpha numeric
	str1 = RandomAlphanumeric(21)
	assertEquals(t, len(str1), 21, "RandomAlphanumeric(21)")
	assertTrue(t, func() bool {
		alphanumeric, _ := regexp.MatchString("^[0-9a-zA-Z]+$", str1)
		return alphanumeric
	}(), "PASS", "str1 contains alpha or numeric")
	str2 = RandomAlphanumeric(21)
	assertTrue(t, str1 != str2, "PASS", "str1 != str2")
	fmt.Printf("RandomAlphanumeric(21):\nstr1=%s\nstr2=%s\n", str1, str2)

	// random specified chars
	strSet := []rune("囧ABCxyz")
	str1 = RandomStringSpec0(21, strSet)
	assertEquals(t, len([]rune(str1)), 21, "RandomSpec0(21) length")
	assertTrue(t, func() bool {
		match, _ := regexp.MatchString("^[囧ABCxyz]+$", str1)
		return match
	}(), "PASS", fmt.Sprintf("Only contains %s", "囧ABCxyz"))
	str2 = RandomStringSpec0(21, strSet)
	assertTrue(t, str1 != str2, "PASS", "str1 != str2")
	fmt.Printf("RandomSpec0(21):\nstr1=%s\nstr2=%s\n", str1, str2)
}
