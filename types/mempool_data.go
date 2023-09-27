package types

type PoolTransactionData struct {
	BlockHash            string         `json:"blockHash"`
	BlockNumber          string         `json:"blockNumber"`
	From                 string         `json:"from"`
	Gas                  string         `json:"gas"`
	GasPrice             string         `json:"gasPrice"`
	MaxFeePerGas         string         `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string         `json:"maxPriorityFeePerGas"`
	Hash                 string         `json:"hash"`
	Input                string         `json:"input"`
	Nonce                string         `json:"nonce"`
	To                   string         `json:"to"`
	TransactionIndex     interface{}    `json:"transactionIndex"`
	Value                string         `json:"value"`
	Type                 string         `json:"type"`
	AccessList           []interface{}  `json:"accessList"`
	ChainID              string         `json:"chainId"`
	V                    string         `json:"v"`
	R                    string         `json:"r"`
	S                    string         `json:"s"`
	Creates              interface{}    `json:"creates"`
	Wait                 interface{}    `json:"wait"`
}

type MemPool struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result ResultData `json:"result"`
}

type ResultData struct {
	
		Pending map[string]map[string]PoolTransactionData `json:"pending"`
		Queued  map[string]map[string]PoolTransactionData `json:"queued"`
	
}