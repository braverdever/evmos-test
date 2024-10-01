package models

type Block struct {
	TotalDifficulty  string        `json:"totalDifficulty"`
	Uncles           []any         `json:"uncles"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	ParentHash       string        `json:"parentHash"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Difficulty       string        `json:"difficulty"`
	LogsBloom        string        `json:"logsBloom"`
	Number           string        `json:"number"`
	Size             string        `json:"size"`
	Nonce            string        `json:"nonce"`
	Timestamp        string        `json:"timestamp"`
	TransactionsRoot string        `json:"transactionsRoot"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	StateRoot        string        `json:"stateRoot"`
	Transactions     []Transaction `json:"transactions"`
}

type ResponseBlock struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  Block  `json:"result"`
}
