package main

import (
	"ethernal/explorer-api/docs"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @BasePath /api

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

// @BasePath /api

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example2
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
// @BasePath /
func main() {
	fmt.Println("The number of CPU Cores:", runtime.NumCPU())
	server := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	v1 := server.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}

	v2 := server.Group("/api/v2")
	{
		eg := v2.Group("/example")
		{
			eg.GET("/helloworld", Helloworldv2)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run("localhost:8888")
}
