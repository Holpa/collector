package graph

import (
	"math/big"
	"strconv"
)

func ParseInt(value string) int {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return int(parsed)
}

func ParseBigFloat(value string) *big.Float {
	b := big.NewFloat(0)
	v, ok := b.SetString(value)
	if !ok {
		return big.NewFloat(0)
	}
	return v
}

func ParseUInt(value string) uint {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}
	return uint(parsed)
}
