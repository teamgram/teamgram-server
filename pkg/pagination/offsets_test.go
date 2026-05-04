package pagination

import "testing"

func TestNormalizeLimit(t *testing.T) {
	tests := []struct {
		name  string
		limit int32
		want  int32
	}{
		{name: "default", limit: 0, want: DefaultLimit},
		{name: "negative default", limit: -1, want: DefaultLimit},
		{name: "kept", limit: 50, want: 50},
		{name: "capped", limit: 500, want: MaxLimit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeLimit(tt.limit); got != tt.want {
				t.Fatalf("NormalizeLimit(%d) = %d, want %d", tt.limit, got, tt.want)
			}
		})
	}
}

func TestSliceOffset(t *testing.T) {
	tests := []struct {
		name         string
		offsetID     int32
		offsetFromID int64
		addOffset    int32
		want         int64
	}{
		{name: "initial", want: 0},
		{name: "older than id", offsetID: 50, offsetFromID: 51, want: 51},
		{name: "newer than id", offsetID: 50, offsetFromID: 51, addOffset: -20, want: 31},
		{name: "around id clamps", offsetID: 1, offsetFromID: 100, addOffset: -200, want: 0},
		{name: "plain positive add offset", addOffset: 10, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceOffset(tt.offsetFromID, IDOffsetInput{
				OffsetID:  tt.offsetID,
				AddOffset: tt.addOffset,
			})
			if got != tt.want {
				t.Fatalf("SliceOffset() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestHashInt64IDs(t *testing.T) {
	ids := []int64{1, 2, 3}
	var want uint64
	for _, id := range ids {
		want ^= want >> 21
		want ^= want << 35
		want ^= want >> 4
		want += uint64(id)
	}
	if got := HashInt64IDs(ids); got != int64(want) {
		t.Fatalf("HashInt64IDs() = %d, want %d", got, int64(want))
	}
}
