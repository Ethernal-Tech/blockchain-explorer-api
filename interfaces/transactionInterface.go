package interfaces

import "ethernal/explorer-api/models"

type TransactionService interface {
	GetTransactionByHash(string) (*models.Transaction, error)
	GetTransactionsInBlock(uint64) ([]models.Transaction, error)
	GetTransactionsByAddress(string, *models.PaginationData) ([]models.Transaction, error)
}
