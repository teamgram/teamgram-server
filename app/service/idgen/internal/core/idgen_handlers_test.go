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
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/repository/alloc"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/svc"
)

type coreSeqStore struct {
	maxSeq         int64
	mallocCalls    []coreSeqMallocCall
	setMaxSeqCalls []coreSeqSetMaxSeqCall
}

type coreSeqMallocCall struct {
	table string
	id    int64
	size  int64
}

type coreSeqSetMaxSeqCall struct {
	table string
	id    int64
	seq   int64
}

func (s *coreSeqStore) Malloc(_ context.Context, table string, id int64, size int64) (int64, error) {
	s.mallocCalls = append(s.mallocCalls, coreSeqMallocCall{table: table, id: id, size: size})
	start := s.maxSeq
	s.maxSeq += size
	return start, nil
}

func (s *coreSeqStore) GetMaxSeq(_ context.Context, _ string, _ int64) (int64, error) {
	return s.maxSeq, nil
}

func (s *coreSeqStore) SetMaxSeq(_ context.Context, table string, id int64, seq int64) error {
	s.setMaxSeqCalls = append(s.setMaxSeqCalls, coreSeqSetMaxSeqCall{table: table, id: id, seq: seq})
	if seq > s.maxSeq {
		s.maxSeq = seq
	}
	return nil
}

func newTestCore(t *testing.T, store *coreSeqStore) *IdgenCore {
	t.Helper()

	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatalf("snowflake.NewNode() err = %v", err)
	}

	var seqAlloc *alloc.Allocator
	if store != nil {
		seqAlloc = alloc.NewAllocator(nil, store)
	}

	return New(context.Background(), &svc.ServiceContext{
		Repo: repository.New(node, seqAlloc),
	})
}

func TestIdgenNextIdsRejectsOutOfRangeNum(t *testing.T) {
	c := newTestCore(t, &coreSeqStore{})

	if _, err := c.IdgenNextIds(&idgen.TLIdgenNextIds{Num: 101}); !errors.Is(err, idgen.ErrInvalidArgument) {
		t.Fatalf("IdgenNextIds() err = %v, want ErrInvalidArgument", err)
	}
	if _, err := c.IdgenNextIds(&idgen.TLIdgenNextIds{Num: -1}); !errors.Is(err, idgen.ErrInvalidArgument) {
		t.Fatalf("IdgenNextIds() err = %v, want ErrInvalidArgument", err)
	}
}

func TestIdgenSeqHandlersUseSeqAllocAndReturnCurrentValue(t *testing.T) {
	store := &coreSeqStore{maxSeq: 10}
	c := newTestCore(t, store)

	next, err := c.IdgenGetNextNSeqId(&idgen.TLIdgenGetNextNSeqId{
		Key: "message_box_ngen_42",
		N:   3,
	})
	if err != nil {
		t.Fatalf("IdgenGetNextNSeqId() err = %v", err)
	}
	if next.V != 13 {
		t.Fatalf("IdgenGetNextNSeqId() = %d, want 13", next.V)
	}
	if len(store.mallocCalls) != 1 {
		t.Fatalf("Malloc calls = %d, want 1", len(store.mallocCalls))
	}
	if got := store.mallocCalls[0]; got != (coreSeqMallocCall{table: alloc.MessageBoxNGen, id: 42, size: 3}) {
		t.Fatalf("Malloc call = %+v, want message_box_ngen id 42 size 3", got)
	}

	current, err := c.IdgenGetCurrentSeqId(&idgen.TLIdgenGetCurrentSeqId{Key: "message_box_ngen_42"})
	if err != nil {
		t.Fatalf("IdgenGetCurrentSeqId() err = %v", err)
	}
	if current.V != 13 {
		t.Fatalf("IdgenGetCurrentSeqId() = %d, want 13", current.V)
	}

	if _, err := c.IdgenSetCurrentSeqId(&idgen.TLIdgenSetCurrentSeqId{Key: "message_box_ngen_42", Id: 20}); err != nil {
		t.Fatalf("IdgenSetCurrentSeqId() err = %v", err)
	}
	if len(store.setMaxSeqCalls) != 1 {
		t.Fatalf("SetMaxSeq calls = %d, want 1", len(store.setMaxSeqCalls))
	}
	if got := store.setMaxSeqCalls[0]; got != (coreSeqSetMaxSeqCall{table: alloc.MessageBoxNGen, id: 42, seq: 20}) {
		t.Fatalf("SetMaxSeq call = %+v, want message_box_ngen id 42 seq 20", got)
	}
}

func TestIdgenGetNextIdValListHandlesSnowflakeAndSeqInputs(t *testing.T) {
	store := &coreSeqStore{maxSeq: 20}
	c := newTestCore(t, store)

	r, err := c.IdgenGetNextIdValList(&idgen.TLIdgenGetNextIdValList{
		Id: []idgen.InputIdClazz{
			idgen.MakeTLInputId(&idgen.TLInputId{}),
			idgen.MakeTLInputIds(&idgen.TLInputIds{Num: 2}),
			idgen.MakeTLInputSeqId(&idgen.TLInputSeqId{Key: "pts_updates_ngen_7"}),
			idgen.MakeTLInputNSeqId(&idgen.TLInputNSeqId{Key: "channel_pts_updates_ngen_8", N: 4}),
		},
	})
	if err != nil {
		t.Fatalf("IdgenGetNextIdValList() err = %v", err)
	}
	if len(r.Datas) != 4 {
		t.Fatalf("len(Datas) = %d, want 4", len(r.Datas))
	}
	if v, ok := r.Datas[0].(*idgen.TLIdVal); !ok || v.Id == 0 {
		t.Fatalf("Datas[0] = %#v, want non-zero idVal", r.Datas[0])
	}
	if v, ok := r.Datas[1].(*idgen.TLIdVals); !ok || len(v.Id) != 2 || v.Id[0] == 0 || v.Id[1] == 0 {
		t.Fatalf("Datas[1] = %#v, want idVals with 2 non-zero ids", r.Datas[1])
	}
	if v, ok := r.Datas[2].(*idgen.TLSeqIdVal); !ok || v.Id != 21 {
		t.Fatalf("Datas[2] = %#v, want seqIdVal 21", r.Datas[2])
	}
	if v, ok := r.Datas[3].(*idgen.TLSeqIdVal); !ok || v.Id != 25 {
		t.Fatalf("Datas[3] = %#v, want seqIdVal 25", r.Datas[3])
	}
}

func TestIdgenGetCurrentSeqIdListRejectsNonSeqInput(t *testing.T) {
	c := newTestCore(t, &coreSeqStore{})

	_, err := c.IdgenGetCurrentSeqIdList(&idgen.TLIdgenGetCurrentSeqIdList{
		Id: []idgen.InputIdClazz{idgen.MakeTLInputId(&idgen.TLInputId{})},
	})
	if !errors.Is(err, idgen.ErrInvalidArgument) {
		t.Fatalf("IdgenGetCurrentSeqIdList() err = %v, want ErrInvalidArgument", err)
	}
}

// TestIdgenSeqHandlersReturnExportedErrors covers the two service-contract
// branches BFF callers are expected to handle: a nil seq allocator (idgen
// is wired in DB-direct mode without a backing store) surfaces as
// ErrSeqAllocatorUnavailable, and a malformed key surfaces as
// ErrInvalidArgument. Both are produced by Repository now, so the test
// just verifies the wiring all the way to the handler.
func TestIdgenSeqHandlersReturnExportedErrors(t *testing.T) {
	c := newTestCore(t, nil)

	_, err := c.IdgenGetNextSeqId(&idgen.TLIdgenGetNextSeqId{Key: "message_box_ngen_42"})
	if !errors.Is(err, idgen.ErrSeqAllocatorUnavailable) {
		t.Fatalf("IdgenGetNextSeqId() err = %v, want ErrSeqAllocatorUnavailable", err)
	}

	c = newTestCore(t, &coreSeqStore{})
	_, err = c.IdgenGetNextSeqId(&idgen.TLIdgenGetNextSeqId{Key: "unknown_42"})
	if !errors.Is(err, idgen.ErrInvalidArgument) {
		t.Fatalf("IdgenGetNextSeqId() err = %v, want ErrInvalidArgument", err)
	}
}

// TestIdgenGetNextNSeqIdRejectsNonPositiveN: the per-call seq batch size
// must be > 0 — n == 0 has no meaningful semantics ("read current" is
// idgen.getCurrentSeqId), and n < 0 has never been valid. Both surface
// idgenpb.ErrInvalidArgument so BFF callers can branch consistently.
func TestIdgenGetNextNSeqIdRejectsNonPositiveN(t *testing.T) {
	c := newTestCore(t, &coreSeqStore{})

	_, err := c.IdgenGetNextNSeqId(&idgen.TLIdgenGetNextNSeqId{Key: "message_box_ngen_42", N: 0})
	if !errors.Is(err, idgen.ErrInvalidArgument) {
		t.Fatalf("IdgenGetNextNSeqId(n=0) err = %v, want ErrInvalidArgument", err)
	}

	_, err = c.IdgenGetNextNSeqId(&idgen.TLIdgenGetNextNSeqId{Key: "message_box_ngen_42", N: -1})
	if !errors.Is(err, idgen.ErrInvalidArgument) {
		t.Fatalf("IdgenGetNextNSeqId(n=-1) err = %v, want ErrInvalidArgument", err)
	}
}
