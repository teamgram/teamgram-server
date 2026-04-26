package repository

import (
	cryptorand "crypto/rand"
	"math"
	"math/big"
)

func secureRandInt63() (int64, error) {
	v, err := cryptorand.Int(cryptorand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}
