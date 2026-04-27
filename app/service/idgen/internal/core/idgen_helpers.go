package core

import (
	"fmt"
	"strconv"
	"strings"

	idgenpb "github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/repository/alloc"
)

const maxNextIdsNum = 100

var (
	// TODO: 将表名白名单可配置化.
	seqTables = []string{
		alloc.MessageDataNGen,
		alloc.MessageBoxNGen,
		alloc.ChannelMessageBoxNGen,
		alloc.SeqUpdatesNGen,
		alloc.PtsUpdatesNGen,
		alloc.QtsUpdatesNGen,
		alloc.ChannelPtsUpdatesNGen,
		alloc.ScheduledNGen,
		alloc.BotUpdatesNGen,
		alloc.StoryNGen,
		alloc.ChannelStoryNGen,
	}
)

func validateNextIdsNum(num int32) error {
	if num < 0 || num > maxNextIdsNum {
		return fmt.Errorf("%w: next ids num %d out of range [0,%d]", idgenpb.ErrInvalidArgument, num, maxNextIdsNum)
	}
	return nil
}

func parseSeqKey(key string) (string, int64, error) {
	for _, table := range seqTables {
		prefix := table + "_"
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		id, err := strconv.ParseInt(strings.TrimPrefix(key, prefix), 10, 64)
		if err != nil {
			return "", 0, fmt.Errorf("%w: parse seq key %q: %v", idgenpb.ErrInvalidArgument, key, err)
		}
		return table, id, nil
	}
	return "", 0, fmt.Errorf("%w: invalid seq key %q", idgenpb.ErrInvalidArgument, key)
}

func (c *IdgenCore) getCurrentSeqID(key string) (int64, error) {
	if c.svcCtx.Repo.SeqAlloc == nil {
		return 0, idgenpb.ErrSeqAllocatorUnavailable
	}
	table, id, err := parseSeqKey(key)
	if err != nil {
		return 0, err
	}
	seq, err := c.svcCtx.Repo.SeqAlloc.GetMaxSeq(c.ctx, table, id)
	if err != nil {
		return 0, fmt.Errorf("%w: get max seq: %w", idgenpb.ErrSeqStorage, err)
	}
	return seq, nil
}

func (c *IdgenCore) setCurrentSeqID(key string, seq int64) error {
	if c.svcCtx.Repo.SeqAlloc == nil {
		return idgenpb.ErrSeqAllocatorUnavailable
	}
	table, id, err := parseSeqKey(key)
	if err != nil {
		return err
	}
	if err := c.svcCtx.Repo.SeqAlloc.SetMaxSeq(c.ctx, table, id, seq); err != nil {
		return fmt.Errorf("%w: set max seq: %w", idgenpb.ErrSeqStorage, err)
	}
	return nil
}

func (c *IdgenCore) getNextSeqID(key string, n int32) (int64, error) {
	if c.svcCtx.Repo.SeqAlloc == nil {
		return 0, idgenpb.ErrSeqAllocatorUnavailable
	}
	if n < 0 {
		return 0, fmt.Errorf("%w: seq n %d must be >= 0", idgenpb.ErrInvalidArgument, n)
	}
	table, id, err := parseSeqKey(key)
	if err != nil {
		return 0, err
	}
	start, err := c.svcCtx.Repo.SeqAlloc.Malloc(c.ctx, table, id, int64(n))
	if err != nil {
		return 0, fmt.Errorf("%w: malloc seq: %w", idgenpb.ErrSeqStorage, err)
	}
	return start + int64(n), nil
}

func (c *IdgenCore) nextID() int64 {
	return c.svcCtx.Repo.Node.Generate().Int64()
}

func (c *IdgenCore) nextIDs(num int32) ([]int64, error) {
	if err := validateNextIdsNum(num); err != nil {
		return nil, err
	}
	ids := make([]int64, num)
	for i := int32(0); i < num; i++ {
		ids[i] = c.nextID()
	}
	return ids, nil
}
