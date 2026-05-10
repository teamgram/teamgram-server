package objectstore

import (
	"crypto/sha256"
	"reflect"
	"testing"
)

func TestBuildHashChunksUsesFixedRanges(t *testing.T) {
	got := BuildHashChunks([]byte("abcdefghij"), 4)
	if len(got) != 3 {
		t.Fatalf("BuildHashChunks() len = %d, want 3", len(got))
	}
	wantLimits := []int32{4, 4, 2}
	for i, chunk := range got {
		if chunk.Offset != int64(i*4) {
			t.Fatalf("chunk %d offset = %d", i, chunk.Offset)
		}
		if chunk.Limit != wantLimits[i] {
			t.Fatalf("chunk %d limit = %d, want %d", i, chunk.Limit, wantLimits[i])
		}
		start := i * 4
		end := start + int(chunk.Limit)
		sum := sha256.Sum256([]byte("abcdefghij")[start:end])
		if !reflect.DeepEqual(chunk.Hash, sum[:]) {
			t.Fatalf("chunk %d hash mismatch", i)
		}
	}
}

func TestFilterHashChunksIntersectsRange(t *testing.T) {
	chunks := []HashChunk{
		{Offset: 0, Limit: 4, Hash: []byte("a")},
		{Offset: 4, Limit: 4, Hash: []byte("b")},
		{Offset: 8, Limit: 2, Hash: []byte("c")},
	}

	got := FilterHashChunks(chunks, 3, 5)
	if len(got) != 2 || got[0].Offset != 0 || got[1].Offset != 4 {
		t.Fatalf("FilterHashChunks() = %#v, want first two intersecting chunks", got)
	}

	got = FilterHashChunks(chunks, 4, 0)
	if len(got) != 2 || got[0].Offset != 4 || got[1].Offset != 8 {
		t.Fatalf("FilterHashChunks() with no limit = %#v, want chunks at/after offset", got)
	}
}
