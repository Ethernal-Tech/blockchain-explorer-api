package database

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
