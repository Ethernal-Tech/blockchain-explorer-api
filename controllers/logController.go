package controllers

import (
	"ethernal/explorer-api/common"
	"ethernal/explorer-api/configuration"
	"ethernal/explorer-api/interfaces"
	"ethernal/explorer-api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	LogService interfaces.LogService
	config     *configuration.Configuration
}

func NewLogController(logService interfaces.LogService, config *configuration.Configuration) LogController {
	return LogController{LogService: logService, config: config}
}

// @Summary Get logs by transaction hash
// @Tags log
// @Produce json
// @Param transactionhash path string true "transaction hash"
// @Success 200 {array} models.Log
// @Failure 500 {string} common.ErrInternal
// @Router /v1/log/transactionhash/{transactionhash} [get]
func (lc *LogController) GetLogsByTransactionHash(context *gin.Context) {
	hash := context.Param("transactionhash")

	logs, err := lc.LogService.GetLogsByTransactionHash(hash)

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

	context.JSON(http.StatusOK, logs)
}

// @Summary Get logs
// @Tags log
// @Produce json
// @Param address query string false "address"
// @Param startBlock query integer false "block number to start searching for logs"
// @Param endBlock query integer false "block number to stop searching for logs"
// @Param topic0 query string false "topic0"
// @Param topic1 query string false "topic1"
// @Param topic2 query string false "topic2"
// @Param topic3 query string false "topic3"
// @Param topic0_1_opr query string false "operator and|or between topic0 & topic1"
// @Param topic0_2_opr query string false "operator and|or between topic0 & topic2"
// @Param topic0_3_opr query string false "operator and|or between topic0 & topic3"
// @Param topic1_2_opr query string false "operator and|or between topic1 & topic2"
// @Param topic1_3_opr query string false "operator and|or between topic1 & topic3"
// @Param topic2_3_opr query string false "operator and|or between topic2 & topic3"
// @Param page query integer false "page number"
// @Param perPage query integer false "number of logs displayed per page"
// @Param sort query string false "use asc to sort by ascending and desc to sort by descending"
// @Success 200 {array} models.Log
// @Failure 400 {string} common.BadPaginationParams
// @Failure 500 {string} common.ErrInternal
// @Router /v1/log [get]
func (lc *LogController) GetLogs(context *gin.Context) {
	address := context.Query("address")
	topicData := getTopicData(context)

	if address == "" && topicData.Topic0 == "" && topicData.Topic1 == "" && topicData.Topic2 == "" && topicData.Topic3 == "" {
		context.JSON(http.StatusBadRequest, common.AddressOrTopicRequired)
		return
	}

	paginationData := GetPaginationData(context, lc.config.PaginationMaxRecords, lc.config.PaginationSort)

	if paginationData.Page*paginationData.PerPage > int(lc.config.PaginationMaxRecords) {
		context.JSON(http.StatusBadRequest, fmt.Sprint(common.BadPaginationParams, lc.config.PaginationMaxRecords))
		return
	}

	logs, err := lc.LogService.GetLogs(address, topicData, paginationData)

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

	context.JSON(http.StatusOK, logs)
}

func getTopicData(context *gin.Context) *models.TopicData {
	var topicData = models.TopicData{
		Topic0:       strings.ToLower(context.Query("topic0")),
		Topic1:       strings.ToLower(context.Query("topic1")),
		Topic2:       strings.ToLower(context.Query("topic2")),
		Topic3:       strings.ToLower(context.Query("topic3")),
		Topic0_1_opr: common.And,
		Topic0_2_opr: common.And,
		Topic0_3_opr: common.And,
		Topic1_2_opr: common.And,
		Topic1_3_opr: common.And,
		Topic2_3_opr: common.And,
	}

	topic0_1_opr := context.Query("topic0_1_opr")
	if strings.ToLower(topic0_1_opr) == common.Or {
		topicData.Topic0_1_opr = common.Or
	}

	topic0_2_opr := context.Query("topic0_2_opr")
	if strings.ToLower(topic0_2_opr) == common.Or {
		topicData.Topic0_2_opr = common.Or
	}

	topic0_3_opr := context.Query("topic0_3_opr")
	if strings.ToLower(topic0_3_opr) == common.Or {
		topicData.Topic0_3_opr = common.Or
	}

	topic1_2_opr := context.Query("topic1_2_opr")
	if strings.ToLower(topic1_2_opr) == common.Or {
		topicData.Topic1_2_opr = common.Or
	}

	topic1_3_opr := context.Query("topic1_3_opr")
	if strings.ToLower(topic1_3_opr) == common.Or {
		topicData.Topic1_3_opr = common.Or
	}

	topic2_3_opr := context.Query("topic2_3_opr")
	if strings.ToLower(topic2_3_opr) == common.Or {
		topicData.Topic2_3_opr = common.Or
	}

	return &topicData
}
