package models

import "math/big"

type RequestBlock struct {
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
}

type Contract struct {
	Address string
	Count   int
}

type Wallet struct {
	Address string
	Balance *big.Float
}
