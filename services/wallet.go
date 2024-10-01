package services

import (
	"encoding/json"
	"evmos-test/models"
	"evmos-test/pkgs/functions"
	"math/big"
)

func GetBalance(hash string, url string) (*big.Float, error) {
	request := models.RequestBlock{
		Method: "eth_getBalance",
		Params: []any{
			hash,
			"latest",
		},
		Id:      1,
		JsonRpc: "2.0",
	}
	data, err := Post(url, request)
	if err != nil {
		return nil, err
	}

	response := models.ResponseBalance{}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return functions.FloatFromHex(response.Result), nil
}
