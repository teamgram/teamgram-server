// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package xkv

import "testing"

func TestComputeSaltsTTLUsesMaxValidUntil(t *testing.T) {
	now := int32(1_700_000_000)
	salts := []*FutureSaltRecord{
		{ValidSince: now, ValidUntil: now + 1800, Salt: 1},
		{ValidSince: now + 1800, ValidUntil: now + 7200, Salt: 2},
	}

	got := computeSaltsTTL(salts, now)
	want := 7200 + saltGraceTTL
	if got != want {
		t.Fatalf("computeSaltsTTL = %d, want %d", got, want)
	}
}

func TestComputeSaltsTTLEnforcesMinimum(t *testing.T) {
	now := int32(1_700_000_000)
	salts := []*FutureSaltRecord{
		{ValidSince: now, ValidUntil: now - 600, Salt: 99},
	}

	if got := computeSaltsTTL(salts, now); got != saltMinTTL {
		t.Fatalf("computeSaltsTTL = %d, want at least %d", got, saltMinTTL)
	}
}

func TestComputeSaltsTTLEmpty(t *testing.T) {
	if got := computeSaltsTTL(nil, 1_700_000_000); got != 0 {
		t.Fatalf("computeSaltsTTL(nil) = %d, want 0", got)
	}
}
