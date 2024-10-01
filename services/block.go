package services

import (
	"encoding/json"
	"evmos-test/models"
	"fmt"
)

func GetBlock(id int, url string) (*models.Block, error) {
	hexOfId := fmt.Sprintf("0x%x", id)
	request := models.RequestBlock{
		Method: "eth_getBlockByNumber",
		Params: []any{
			hexOfId,
			true,
		},
		Id:      1,
		JsonRpc: "2.0",
	}
	data, err := Post(url, request)
	if err != nil {
		return nil, err
	}

	response := models.ResponseBlock{}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}
