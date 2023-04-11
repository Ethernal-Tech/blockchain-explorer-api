package controllers

import (
	"ethernal/explorer-api/common"
	"ethernal/explorer-api/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlockController struct {
	BlockService interfaces.BlockService
}

func NewBlockController(blockService interfaces.BlockService) BlockController {
	return BlockController{BlockService: blockService}
}

// @Summary Get block by number
// @Tags block
// @Produce json
// @Param blocknumber path uint64 true "block number"
// @Success 200 {object} models.Block
// @Failure 400 {string} common.BadBlockNumber
// @Failure 404 {string} common.ErrNotFound
// @Failure 500 {string} common.ErrInternal
// @Router /v1/block/number/{blocknumber} [get]
func (bc *BlockController) GetBlockByNumber(gin *gin.Context) {
	num := gin.Param("blocknumber")
	blockNumber, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		gin.JSON(http.StatusBadRequest, common.BadBlockNumber)
		return
	}
	block, err := bc.BlockService.GetBlockByNumber(blockNumber)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			gin.JSON(http.StatusNotFound, common.ErrNotFound.Error())
			return
		default:
			gin.JSON(http.StatusInternalServerError, common.ErrInternal.Error())
			return
		}
	}
	gin.JSON(http.StatusOK, block)
}

// @Summary Get block by hash
// @Tags block
// @Produce json
// @Param blockhash path string true "block hash"
// @Success 200 {object} models.Block
// @Failure 404 {string} common.ErrNotFound
// @Failure 500 {string} common.ErrInternal
// @Router /v1/block/hash/{blockhash} [get]
func (bc *BlockController) GetBlockByHash(gin *gin.Context) {
	blockHash := gin.Param("blockhash")

	block, err := bc.BlockService.GetBlockByHash(blockHash)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			gin.JSON(http.StatusNotFound, common.ErrNotFound.Error())
			return
		default:
			gin.JSON(http.StatusInternalServerError, common.ErrInternal.Error())
			return
		}
	}
	gin.JSON(http.StatusOK, block)
}
