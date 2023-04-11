package models

type Transaction struct {
	Hash             string `json:"hash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      uint64 `json:"blockNumber"`
	From             string `json:"from"`
	To               string `json:"to"`
	Gas              uint64 `json:"gas"`
	GasUsed          uint64 `json:"gasUsed"`
	GasPrice         uint64 `json:"gasPrice"`
	Nonce            uint64 `json:"nonce"`
	TransactionIndex uint64 `json:"transactionIndex"`
	Value            string `json:"value"`
	ContractAddress  string `json:"contractAddress"`
	Status           uint64 `json:"status"`
	Timestamp        uint64 `json:"timestamp"`
	InputData        string `json:"input"`
}
