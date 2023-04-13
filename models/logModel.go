package models

type Log struct {
	BlockHash       string   `json:"blockHash"`
	LogIndex        uint32   `json:"logIndex"`
	TransactionHash string   `json:"transactionHash"`
	Address         string   `json:"address"`
	BlockNumber     uint64   `json:"blockNumber"`
	Topics          []string `json:"topics"`
	Data            string   `json:"data"`
}
