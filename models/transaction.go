package models

type Transaction struct {
	Hash             string `json:"hash"`
	Value            string `json:"value"`
	S                string `json:"s"`
	Gas              string `json:"gas"`
	Nonce            string `json:"nonce"`
	Type             string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	BlockNumber      string `json:"blockNumber"`
	Input            string `json:"input"`
	TransactionIndex string `json:"transactionIndex"`
	BlockHash        string `json:"blockHash"`
	GasPrice         string `json:"gasPrice"`
	To               string `json:"to"`
	ChainId          string `json:"chainId"`
	From             string `json:"from"`
}

type TransactionReceipt struct {
	Logs              []interface{} `json:"logs"`
	To                string        `json:"to"`
	TransactionHash   string        `json:"transactionHash"`
	BlockHash         string        `json:"blockHash"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed"`
	From              string        `json:"from"`
	LogsBloom         string        `json:"logsBloom"`
	Status            string        `json:"status"`
	TransactionIndex  string        `json:"transactionIndex"`
	Type              string        `json:"type"`
	BlockNumber       string        `json:"blockNumber"`
	ContractAddress   any           `json:"contractAddress"`
	GasUsed           string        `json:"gasUsed"`
}

type TransactionTrace struct {
	Gas     string `json:"gas"`
	GasUsed string `json:"gasUsed"`
	Input   string `json:"input"`
	Output  string `json:"output"`
	To      string `json:"to"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	From    string `json:"from"`
}

type ResponseTransactionReceipt struct {
	JsonRpc string             `json:"jsonrpc"`
	Id      int                `json:"id"`
	Result  TransactionReceipt `json:"result"`
}

type ResponseTransactionTrace struct {
	JsonRpc string           `json:"jsonrpc"`
	Id      int              `json:"id"`
	Result  TransactionTrace `json:"result"`
}
