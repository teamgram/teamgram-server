package tg

import (
	"fmt"
	"math"
)

func DateInt32FromUnixSeconds(seconds int64) (int32, error) {
	if seconds < 0 || seconds > math.MaxInt32 {
		return 0, fmt.Errorf("tg date out of int32 range: %d", seconds)
	}
	return int32(seconds), nil
}
