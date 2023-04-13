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

type LogService struct {
	database *bun.DB
	ctx      context.Context
}

func NewLogService(database *bun.DB, ctx context.Context) interfaces.LogService {
	return &LogService{database: database, ctx: ctx}
}

func (ls *LogService) GetLogsByTransactionHash(hash string) ([]models.Log, error) {
	dbLogs := make([]database.Log, 0)
	hash = strings.ToLower(hash)

	// fetch logs from database
	if err := ls.database.NewSelect().Model((*database.Log)(nil)).Where("transaction_hash = ?", hash).Order("index ASC").Scan(ls.ctx, &dbLogs); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database logs to DTOs
	logs := make([]models.Log, 0)

	for _, dbLog := range dbLogs {
		var log = models.Log{
			BlockHash:       dbLog.BlockHash,
			LogIndex:        dbLog.Index,
			TransactionHash: dbLog.TransactionHash,
			Address:         dbLog.Address,
			BlockNumber:     dbLog.BlockNumber,
			Topics:          []string{},
			Data:            dbLog.Data,
		}
		topics := []string{dbLog.Topic0, dbLog.Topic1, dbLog.Topic2, dbLog.Topic3}
		for _, topic := range topics {
			if topic != "" {
				log.Topics = append(log.Topics, topic)
			}
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func (ls *LogService) GetLogs(address string, topicData *models.TopicData, pagination *models.PaginationData) ([]models.Log, error) {
	var offset = pagination.PerPage * (pagination.Page - 1)
	address = strings.ToLower(address)
	dbLogs := make([]database.Log, 0)

	query := ls.database.NewSelect().Model((*database.Log)(nil))
	// if an address is requested, add it to the query
	if address != "" {
		query.Where("address = ?", address)
	}
	// if topics are requested, add them to the query
	if topicData.Topic0 != "" || topicData.Topic1 != "" || topicData.Topic2 != "" || topicData.Topic3 != "" {
		query.WhereGroup(" AND ", func(query *bun.SelectQuery) *bun.SelectQuery {
			if topicData.Topic0 != "" {
				query.Where("topic0 = ?", topicData.Topic0)
			}
			if topicData.Topic1 != "" {
				if topicData.Topic0 != "" && topicData.Topic0_1_opr == common.Or {
					query.WhereOr("topic1 = ?", topicData.Topic1)
				} else {
					query.Where("topic1 = ?", topicData.Topic1)
				}
			}
			if topicData.Topic2 != "" {
				if (topicData.Topic1 != "" && topicData.Topic1_2_opr == common.Or) || (topicData.Topic0 != "" && topicData.Topic0_2_opr == common.Or) {
					query.WhereOr("topic2 = ?", topicData.Topic2)
				} else {
					query.Where("topic2 = ?", topicData.Topic2)
				}
			}
			if topicData.Topic3 != "" {
				if (topicData.Topic2 != "" && topicData.Topic2_3_opr == common.Or) || (topicData.Topic1 != "" && topicData.Topic1_3_opr == common.Or) || (topicData.Topic0 != "" && topicData.Topic0_3_opr == common.Or) {
					query.WhereOr("topic3 = ?", topicData.Topic2)
				} else {
					query.Where("topic3 = ?", topicData.Topic2)
				}
			}
			return query
		})
	}

	// fetch logs from database
	if err := query.Where("block_number BETWEEN ? AND ?", pagination.StartBlock, pagination.EndBlock).Order("block_number "+pagination.Sort).Order("index ASC").Offset(offset).Limit(pagination.PerPage).Scan(ls.ctx, &dbLogs); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		return nil, common.ErrInternal
	}

	// map database logs to DTOs
	logs := make([]models.Log, 0)

	for _, dbLog := range dbLogs {
		var log = models.Log{
			BlockHash:       dbLog.BlockHash,
			LogIndex:        dbLog.Index,
			TransactionHash: dbLog.TransactionHash,
			Address:         dbLog.Address,
			BlockNumber:     dbLog.BlockNumber,
			Topics:          []string{},
			Data:            dbLog.Data,
		}
		topics := []string{dbLog.Topic0, dbLog.Topic1, dbLog.Topic2, dbLog.Topic3}
		for _, topic := range topics {
			if topic != "" {
				log.Topics = append(log.Topics, topic)
			}
		}

		logs = append(logs, log)
	}

	return logs, nil
}
