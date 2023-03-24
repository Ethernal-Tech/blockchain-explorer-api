package controllers

import (
	"ethernal/explorer-api/common"
	"ethernal/explorer-api/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.TransactionService
}

func NewTransactionController(transactionService interfaces.TransactionService) TransactionController {
	return TransactionController{TransactionService: transactionService}
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
