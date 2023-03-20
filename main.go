package main

import (
	_ "ethernal/explorer-api/docs"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Go + Gin Block Explorer API
// @version 1.0
// @description This is a block explorer server. You can visit the GitHub repository at https://github.com/Ethernal-Tech/blockchain-explorer-api

// @host localhost:8888
// @BasePath /
func main() {
	fmt.Println("The number of CPU Cores:", runtime.NumCPU())
	server := gin.Default()
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run("localhost:8888")
}
