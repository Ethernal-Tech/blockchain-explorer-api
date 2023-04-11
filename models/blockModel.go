package models

type Block struct {
	Hash              string `json:"hash"`
	Number            uint64 `json:"number"`
	ParentHash        string `json:"parentHash"`
	Nonce             string `json:"nonce"`
	Validator         string `json:"validator"`
	Difficulty        string `json:"difficulty"`
	TotalDifficulty   string `json:"totalDifficulty"`
	ExtraData         string `json:"extraData"`
	Size              uint64 `json:"size"`
	GasLimit          uint64 `json:"gasLimit"`
	GasUsed           uint64 `json:"gasUsed"`
	Timestamp         uint64 `json:"timestamp"`
	TransactionsCount int    `json:"transactionsCount"`
}
