package main

import (
	"context"
	"ethernal/explorer-api/configuration"
	"ethernal/explorer-api/controllers"
	"ethernal/explorer-api/database"
	_ "ethernal/explorer-api/docs"
	"ethernal/explorer-api/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	server                *gin.Engine
	transactionController controllers.TransactionController
	blockController       controllers.BlockController
	logController         controllers.LogController
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v2/example/helloworld [get]
func Helloworldv2(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld v2")
}

// @title Block Explorer API
// @version 1.0
// @description This is a block explorer server. You can visit the GitHub repository at https://github.com/Ethernal-Tech/blockchain-explorer-api

// @host localhost:8888
// @BasePath /api
func main() {
	ctx := context.TODO()
	configuration := configuration.Load()
	database := database.Initialize(configuration)
	transactionService := services.NewTransactionService(database, ctx)
	transactionController = controllers.NewTransactionController(transactionService, configuration)
	blockService := services.NewBlockService(database, ctx)
	blockController = controllers.NewBlockController(blockService)
	logService := services.NewLogService(database, ctx)
	logController = controllers.NewLogController(logService, configuration)

	server = gin.Default()

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	server.Use(cors.Default())

	// example of error handling using middleware
	// server = gin.New()
	// server.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler)

	routes()
	// use ginSwagger middleware to serve the API docs
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8888")
}

// routes defines endpoint paths and assigns the handler functions to them.
func routes() {
	v1 := server.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}

		transaction := v1.Group("/transaction")
		{
			transaction.GET("/hash/:txhash", transactionController.GetTransactionByHash)
			transaction.GET("/txinblock/:blocknumber", transactionController.GetTransactionsInBlock)
			transaction.GET("/address/:address", transactionController.GetTransactionsByAddress)
		}

		block := v1.Group("/block")
		{
			block.GET("/number/:blocknumber", blockController.GetBlockByNumber)
			block.GET("/hash/:blockhash", blockController.GetBlockByHash)
		}

		log := v1.Group("/log")
		{
			log.GET("/transactionhash/:transactionhash", logController.GetLogsByTransactionHash)
			log.GET("", logController.GetLogs)
		}
	}

	v2 := server.Group("/api/v2")
	{
		eg := v2.Group("/example")
		{
			eg.GET("/helloworld", Helloworldv2)
		}
	}
}
