package database

type Block struct {
	Hash              string
	Number            uint64
	ParentHash        string
	Nonce             string
	Miner             string
	Difficulty        string
	TotalDifficulty   string
	ExtraData         []byte
	Size              uint64
	GasLimit          uint64
	GasUsed           uint64
	Timestamp         uint64
	TransactionsCount int
}

type Transaction struct {
	Hash             string
	BlockHash        string
	BlockNumber      uint64
	From             string
	To               string
	Gas              uint64
	GasUsed          uint64
	GasPrice         uint64
	Nonce            uint64
	TransactionIndex uint64
	Value            string
	ContractAddress  string
	Status           uint64
	Timestamp        uint64
	InputData        string
}

type Log struct {
	BlockHash       string
	Index           uint32
	TransactionHash string
	Address         string
	BlockNumber     uint64
	Topic1          string
	Topic2          string
	Topic3          string
	Topic4          string
	Data            string
}
