package repository

import "testing"

func TestUnixOrZero(t *testing.T) {
	for _, tc := range []struct {
		name    string
		seconds int64
		want    int64
	}{
		{name: "positive", seconds: 1_778_201_611, want: 1_778_201_611},
		{name: "zero", seconds: 0, want: 0},
		{name: "negative", seconds: -1, want: 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := unixOrZero(tc.seconds); got != tc.want {
				t.Fatalf("unixOrZero(%d) = %d, want %d", tc.seconds, got, tc.want)
			}
		})
	}
}

func TestUnixOptionalString(t *testing.T) {
	for _, tc := range []struct {
		name    string
		seconds int64
		want    string
	}{
		{name: "positive", seconds: 1_778_201_611, want: "1778201611"},
		{name: "zero", seconds: 0, want: ""},
		{name: "negative", seconds: -1, want: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := unixOptionalString(tc.seconds); got != tc.want {
				t.Fatalf("unixOptionalString(%d) = %q, want %q", tc.seconds, got, tc.want)
			}
		})
	}
}
