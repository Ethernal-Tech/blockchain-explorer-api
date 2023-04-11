package interfaces

import "ethernal/explorer-api/models"

type BlockService interface {
	GetBlockByNumber(uint64) (*models.Block, error)
	GetBlockByHash(string) (*models.Block, error)
}
