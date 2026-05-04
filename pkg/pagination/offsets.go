package pagination

const (
	DefaultLimit int32 = 20
	MaxLimit     int32 = 100
)

type IDOffsetInput struct {
	OffsetID  int32
	AddOffset int32
	Limit     int32
}

func NormalizeLimit(limit int32) int32 {
	if limit <= 0 {
		return DefaultLimit
	}
	if limit > MaxLimit {
		return MaxLimit
	}
	return limit
}

func OffsetFromIDPosition(offsetFromID int64, addOffset int32) int64 {
	offset := offsetFromID + int64(addOffset)
	if offset < 0 {
		return 0
	}
	return offset
}

func SliceOffset(offsetFromID int64, in IDOffsetInput) int64 {
	if in.OffsetID <= 0 {
		return OffsetFromIDPosition(0, in.AddOffset)
	}
	return OffsetFromIDPosition(offsetFromID, in.AddOffset)
}

func HashInt64IDs(ids []int64) int64 {
	var h uint64
	for _, id := range ids {
		h ^= h >> 21
		h ^= h << 35
		h ^= h >> 4
		h += uint64(id)
	}
	return int64(h)
}
