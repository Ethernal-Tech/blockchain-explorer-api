package services

import (
	"context"
	"database/sql"
	"errors"
	"ethernal/explorer-api/common"
	"ethernal/explorer-api/database"
	"ethernal/explorer-api/interfaces"
	"ethernal/explorer-api/models"
	"strings"

	"github.com/uptrace/bun"
)

type TransactionService struct {
	database *bun.DB
	ctx      context.Context
}

func NewTransactionService(database *bun.DB, ctx context.Context) interfaces.TransactionService {
	return &TransactionService{database: database, ctx: ctx}
}

// GetTransactionByHash returns the transaction with the given hash.
func (ts *TransactionService) GetTransactionByHash(transactionHash string) (*models.Transaction, error) {
	var dbTransaction database.Transaction
	transactionHash = strings.ToLower(transactionHash)

	// fetch transaction from database
	if err := ts.database.NewSelect().Model(((*database.Transaction)(nil))).Where("hash = ?", transactionHash).Scan(ts.ctx, &dbTransaction); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database transaction to DTO
	var transaction = models.Transaction{
		Hash:             dbTransaction.Hash,
		BlockHash:        dbTransaction.BlockHash,
		BlockNumber:      dbTransaction.BlockNumber,
		From:             dbTransaction.From,
		To:               dbTransaction.To,
		Gas:              dbTransaction.Gas,
		GasUsed:          dbTransaction.GasUsed,
		GasPrice:         dbTransaction.GasPrice,
		Nonce:            dbTransaction.Nonce,
		TransactionIndex: dbTransaction.TransactionIndex,
		Value:            dbTransaction.Value,
		ContractAddress:  dbTransaction.ContractAddress,
		Status:           dbTransaction.Status,
		Timestamp:        dbTransaction.Timestamp,
		InputData:        dbTransaction.InputData,
	}

	return &transaction, nil
}

// GetTransactionsInBlock returns transactions in block with given number.
func (ts *TransactionService) GetTransactionsInBlock(blockNumber uint64) ([]models.Transaction, error) {
	dbTransactions := make([]database.Transaction, 0)

	// fetch transactions from database
	if err := ts.database.NewSelect().Model(((*database.Transaction)(nil))).Where("block_number = ?", blockNumber).Scan(ts.ctx, &dbTransactions); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database transactions to DTOs
	transactions := make([]models.Transaction, 0)

	for _, dbTransaction := range dbTransactions {
		var transaction = models.Transaction{
			Hash:             dbTransaction.Hash,
			BlockHash:        dbTransaction.BlockHash,
			BlockNumber:      dbTransaction.BlockNumber,
			From:             dbTransaction.From,
			To:               dbTransaction.To,
			Gas:              dbTransaction.Gas,
			GasUsed:          dbTransaction.GasUsed,
			GasPrice:         dbTransaction.GasPrice,
			Nonce:            dbTransaction.Nonce,
			TransactionIndex: dbTransaction.TransactionIndex,
			Value:            dbTransaction.Value,
			ContractAddress:  dbTransaction.ContractAddress,
			Status:           dbTransaction.Status,
			Timestamp:        dbTransaction.Timestamp,
			InputData:        dbTransaction.InputData,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (ts *TransactionService) GetTransactionsByAddress(address string, pagination *models.PaginationData) ([]models.Transaction, error) {
	var offset = pagination.PerPage * (pagination.Page - 1)
	address = strings.ToLower(address)
	dbTransactions := make([]database.Transaction, 0)

	// fetch transactions from database
	if err := ts.database.NewSelect().Model((*database.Transaction)(nil)).Where("? = ? OR ? = ?", bun.Ident("from"), address, bun.Ident("to"), address).Where("block_number BETWEEN ? AND ?", pagination.StartBlock, pagination.EndBlock).Order("block_number "+pagination.Sort).Offset(offset).Limit(pagination.PerPage).Scan(ts.ctx, &dbTransactions); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database transactions to DTOs
	transactions := make([]models.Transaction, 0)

	for _, dbTransaction := range dbTransactions {
		var transaction = models.Transaction{
			Hash:             dbTransaction.Hash,
			BlockHash:        dbTransaction.BlockHash,
			BlockNumber:      dbTransaction.BlockNumber,
			From:             dbTransaction.From,
			To:               dbTransaction.To,
			Gas:              dbTransaction.Gas,
			GasUsed:          dbTransaction.GasUsed,
			GasPrice:         dbTransaction.GasPrice,
			Nonce:            dbTransaction.Nonce,
			TransactionIndex: dbTransaction.TransactionIndex,
			Value:            dbTransaction.Value,
			ContractAddress:  dbTransaction.ContractAddress,
			Status:           dbTransaction.Status,
			Timestamp:        dbTransaction.Timestamp,
			InputData:        dbTransaction.InputData,
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
