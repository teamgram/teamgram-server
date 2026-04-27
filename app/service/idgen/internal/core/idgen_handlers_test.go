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

	return New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			Node:     node,
			SeqAlloc: alloc.NewAllocator(nil, store),
		},
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

func TestIdgenSeqHandlersReturnExportedErrors(t *testing.T) {
	c := newTestCore(t, &coreSeqStore{})
	c.svcCtx.Repo.SeqAlloc = nil

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
