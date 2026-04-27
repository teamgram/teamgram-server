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

package alloc

import (
	"errors"
	"testing"
)

func TestQuoteSeqTableAllowsNGenTablesOnly(t *testing.T) {
	tables := []string{
		MessageDataNGen,
		MessageBoxNGen,
		ChannelMessageBoxNGen,
		SeqUpdatesNGen,
		PtsUpdatesNGen,
		QtsUpdatesNGen,
		ChannelPtsUpdatesNGen,
		ScheduledNGen,
		BotUpdatesNGen,
		StoryNGen,
		ChannelStoryNGen,
	}
	for _, table := range tables {
		got, err := quoteSeqTable(table)
		if err != nil {
			t.Fatalf("quoteSeqTable(%q) err = %v", table, err)
		}
		if got != "`"+table+"`" {
			t.Fatalf("quoteSeqTable(%q) = %q, want backtick quoted table", table, got)
		}
	}
}

func TestQuoteSeqTableRejectsUnknownOrUnsafeNames(t *testing.T) {
	for _, table := range []string{"seq_conversations", "message_data_ngen;drop table users", "MessageDataNGen", ""} {
		if _, err := quoteSeqTable(table); !errors.Is(err, ErrInvalidTable) {
			t.Fatalf("quoteSeqTable(%q) err = %v, want ErrInvalidTable", table, err)
		}
	}
}

func TestMySQLQueriesUseParameterizedIDColumn(t *testing.T) {
	table := "`message_box_ngen`"
	tests := map[string]string{
		"get":    getMaxSeqQuery(table),
		"insert": insertMaxSeqQuery(table),
		"ensure": ensureRowQuery(table),
		"lock":   lockMaxSeqQuery(table),
		"update": updateMaxSeqQuery(table),
	}
	want := map[string]string{
		"get":    "select max_seq from `message_box_ngen` where id = ? limit 1",
		"insert": "insert into `message_box_ngen`(id, min_seq, max_seq) values (?, 0, ?) on duplicate key update max_seq = greatest(max_seq, values(max_seq))",
		"ensure": "insert ignore into `message_box_ngen`(id, min_seq, max_seq) values (?, 0, 0)",
		"lock":   "select max_seq from `message_box_ngen` where id = ? for update",
		"update": "update `message_box_ngen` set max_seq = ? where id = ?",
	}
	for name, got := range tests {
		if got != want[name] {
			t.Fatalf("%s query = %q, want %q", name, got, want[name])
		}
	}
}
