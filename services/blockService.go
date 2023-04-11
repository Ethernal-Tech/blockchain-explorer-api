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

type BlockService struct {
	database *bun.DB
	ctx      context.Context
}

func NewBlockService(database *bun.DB, ctx context.Context) interfaces.BlockService {
	return &BlockService{database: database, ctx: ctx}
}

func (bs *BlockService) GetBlockByNumber(number uint64) (*models.Block, error) {
	var dbBlock database.Block

	// fetch block from database
	if err := bs.database.NewSelect().Model((*database.Block)(nil)).Where("number = ?", number).Scan(bs.ctx, &dbBlock); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database block to DTO
	var block = models.Block{
		Hash:              dbBlock.Hash,
		Number:            dbBlock.Number,
		ParentHash:        dbBlock.ParentHash,
		Nonce:             dbBlock.Nonce,
		Validator:         dbBlock.Miner,
		Difficulty:        dbBlock.Difficulty,
		TotalDifficulty:   dbBlock.TotalDifficulty,
		ExtraData:         string(dbBlock.ExtraData),
		Size:              dbBlock.Size,
		GasLimit:          dbBlock.GasLimit,
		GasUsed:           dbBlock.GasUsed,
		Timestamp:         dbBlock.Timestamp,
		TransactionsCount: dbBlock.TransactionsCount,
	}

	return &block, nil
}

func (bs *BlockService) GetBlockByHash(hash string) (*models.Block, error) {
	var dbBlock database.Block
	hash = strings.ToLower(hash)

	// fetch block from database
	if err := bs.database.NewSelect().Model((*database.Block)(nil)).Where("hash = ?", hash).Scan(bs.ctx, &dbBlock); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database block to DTO
	var block = models.Block{
		Hash:              dbBlock.Hash,
		Number:            dbBlock.Number,
		ParentHash:        dbBlock.ParentHash,
		Nonce:             dbBlock.Nonce,
		Validator:         dbBlock.Miner,
		Difficulty:        dbBlock.Difficulty,
		TotalDifficulty:   dbBlock.TotalDifficulty,
		ExtraData:         string(dbBlock.ExtraData),
		Size:              dbBlock.Size,
		GasLimit:          dbBlock.GasLimit,
		GasUsed:           dbBlock.GasUsed,
		Timestamp:         dbBlock.Timestamp,
		TransactionsCount: dbBlock.TransactionsCount,
	}

	return &block, nil
}
