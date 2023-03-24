package middlewares

import (
	"ethernal/explorer-api/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error handling middleware. Middleware serializes error as JSON into the response body.
func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch err.Err {
		case common.ErrNotFound:
			c.JSON(http.StatusNotFound, common.ErrNotFound.Error())
			return
		default:
			c.JSON(http.StatusInternalServerError, common.ErrInternal.Error())
			return
		}
	}
}
