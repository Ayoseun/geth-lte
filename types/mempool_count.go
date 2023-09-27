package types

type MemPoolCount struct {
	JSONRPC string             `json:"jsonrpc"`
	ID      int                `json:"id"`
	Result  MemPoolCountResult `json:"result"`
}

type MemPoolCountResult struct {
	Pending string `json:"pending"`
	Queued  string `json:"queued"`
}
