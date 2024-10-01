package functions

import (
	"math/big"
	"strconv"
)

func FloatFromHex(hex string) *big.Float {
	if len(hex) < 2 {
		return nil
	}
	intValue := new(big.Int)
	intValue.SetString(hex[2:], 16)

	floatValue := new(big.Float).SetInt(intValue)
	floatValue.Quo(floatValue, big.NewFloat(1e9))
	return floatValue
}

func IntFromHex(hex string) int64 {
	if len(hex) < 2 {
		return 0
	}
	value, err := strconv.ParseInt(hex[2:], 16, 64)
	if err != nil {
		return 0
	}
	return value
}
