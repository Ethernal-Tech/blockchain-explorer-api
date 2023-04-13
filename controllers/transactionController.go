package controllers

import (
	"ethernal/explorer-api/common"
	"ethernal/explorer-api/configuration"
	"ethernal/explorer-api/interfaces"
	"ethernal/explorer-api/models"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.TransactionService
	config             *configuration.Configuration
}

func NewTransactionController(transactionService interfaces.TransactionService, config *configuration.Configuration) TransactionController {
	return TransactionController{TransactionService: transactionService, config: config}
}

// @Summary Get transaction by hash
// @Tags transaction
// @Produce json
// @Param txhash path string true "transaction hash"
// @Success 200 {object} models.Transaction
// @Failure 404 {string} common.ErrNotFound
// @Failure 500 {string} common.ErrInternal
// @Router /v1/transaction/hash/{txhash} [get]
func (tc *TransactionController) GetTransactionByHash(context *gin.Context) {
	txhash := context.Param("txhash")

	transaction, err := tc.TransactionService.GetTransactionByHash(txhash)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			context.JSON(http.StatusNotFound, err.Error())
			return
		default:
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	// example of attaching error to current context, which will be processed by middleware for error handling
	// if err != nil {
	// context.Error(err)
	// 	return
	// }

	context.JSON(http.StatusOK, transaction)
}

// @Summary Get transactions in block
// @Tags transaction
// @Produce json
// @Param blocknumber path integer true "block number"
// @Success 200 {array} models.Transaction
// @Failure 400 {string} common.BadBlockNumber
// @Failure 500 {string} common.ErrInternal
// @Router /v1/transaction/txinblock/{blocknumber} [get]
func (tc *TransactionController) GetTransactionsInBlock(context *gin.Context) {
	num := context.Param("blocknumber")
	blockNumber, err := strconv.ParseUint(num, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, common.BadBlockNumber)
		return
	}

	transactions, err := tc.TransactionService.GetTransactionsInBlock(blockNumber)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			context.JSON(http.StatusNotFound, err.Error())
			return
		default:
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	context.JSON(http.StatusOK, transactions)
}

// @Summary Get transactions by address
// @Tags transaction
// @Produce json
// @Param address path string true "address"
// @Param startBlock query integer false "block number to start searching for transactions"
// @Param endBlock query integer false "block number to stop searching for transactions"
// @Param page query integer false "page number"
// @Param perPage query integer false "number of transactions displayed per page"
// @Param sort query string false "use asc to sort by ascending and desc to sort by descending"
// @Success 200 {array} models.Transaction
// @Failure 400 {string} common.BadPaginationParams
// @Failure 500 {string} common.ErrInternal
// @Router /v1/transaction/address/{address} [get]
func (tc *TransactionController) GetTransactionsByAddress(context *gin.Context) {
	address := context.Param("address")

	paginationData := GetPaginationData(context, tc.config.PaginationMaxRecords, tc.config.PaginationSort)

	if paginationData.Page*paginationData.PerPage > int(tc.config.PaginationMaxRecords) {
		context.JSON(http.StatusBadRequest, fmt.Sprint(common.BadPaginationParams, tc.config.PaginationMaxRecords))
		return
	}

	transactions, err := tc.TransactionService.GetTransactionsByAddress(address, paginationData)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			context.JSON(http.StatusNotFound, err.Error())
			return
		default:
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	context.JSON(http.StatusOK, transactions)
}

func GetPaginationData(context *gin.Context, maxRecords int, order string) *models.PaginationData {
	var paginationData = models.PaginationData{
		StartBlock: 0,
		EndBlock:   math.MaxInt64,
		Page:       1,
		PerPage:    maxRecords,
		Sort:       order,
	}

	startBlockStr := context.Query("startBlock")
	if startBlockStr != "" {
		if startBlock, err := strconv.ParseInt(startBlockStr, 10, 64); err == nil {
			paginationData.StartBlock = startBlock
		}
	}

	endBlockStr := context.Query("endBlock")
	if endBlockStr != "" {
		if endBlock, err := strconv.ParseInt(endBlockStr, 10, 64); err == nil {
			paginationData.EndBlock = endBlock
		}
	}

	pageStr := context.Query("page")
	if pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			paginationData.Page = page
		}
	}

	perPageStr := context.Query("perPage")
	if perPageStr != "" {
		if perPage, err := strconv.Atoi(perPageStr); err == nil {
			paginationData.PerPage = perPage
		}
	}

	sort := context.Query("sort")
	if strings.ToLower(sort) == common.Asc {
		paginationData.Sort = sort
	}

	return &paginationData
}
