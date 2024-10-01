package services

import (
	"encoding/json"
	"evmos-test/models"
)

func GetTransactionReceipt(hash string, url string) (*models.TransactionReceipt, error) {
	request := models.RequestBlock{
		Method: "eth_getTransactionReceipt",
		Params: []any{
			hash,
		},
		Id:      1,
		JsonRpc: "2.0",
	}
	data, err := Post(url, request)
	if err != nil {
		return nil, err
	}

	response := models.ResponseTransactionReceipt{}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

func GetTransactionTrace(hash string, tracer string, url string) (*models.TransactionTrace, error) {
	request := models.RequestBlock{
		Method: "debug_traceTransaction",
		Params: []any{
			hash,
			map[string]any{
				"tracer": tracer,
			},
		},
		Id:      1,
		JsonRpc: "2.0",
	}
	data, err := Post(url, request)
	if err != nil {
		return nil, err
	}

	response := models.ResponseTransactionTrace{}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

/*
	rawData := map[string]any{}
	err = json.Unmarshal(data, &rawData)
	if err != nil {
		return err
	}

	for key, value := range rawData {
		fmt.Println(key, reflect.TypeOf(value))
	}

	fmt.Println("-------")

	result := rawData["result"].(map[string]any)
	for key, value := range result {
		fmt.Println(key, reflect.TypeOf(value))
	}

	fmt.Println("-------")
	transactions := result["transactions"].([]any)
	transaction := transactions[0].(map[string]any)

	for key, value := range transaction {
		fmt.Println(key, reflect.TypeOf(value))
	}
*/
