package interfaces

import "ethernal/explorer-api/models"

type LogService interface {
	GetLogsByTransactionHash(string) ([]models.Log, error)
	GetLogs(string, *models.TopicData, *models.PaginationData) ([]models.Log, error)
}
