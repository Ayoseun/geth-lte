package results

type MempoolResult struct {
	BlockHash        *string `json:"blockHash"`
	BlockNumber      *string `json:"blockNumber"`
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	To               string  `json:"to"`
	TransactionIndex *string `json:"transactionIndex"`
	Value            string  `json:"value"`
	Type             string  `json:"type"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
}

type MempoolParams struct {
	Result MempoolResult `json:"result"`
}

type MempoolData struct {
	JSONRPC     string              `json:"jsonrpc"`
	Method      string              `json:"method"`
	Params      MempoolParams `json:"params"`
	Subscription string              `json:"subscription"`
}
