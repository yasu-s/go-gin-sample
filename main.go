package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sample API
// @version 1.0
// @description APIサンプルです
// @host localhost:8080
// @BasePath /api/v1

// API test
// @Summary API test
// @Description APIサンプルです
// @Accept  json
// @Produce  json
// @Router /ping [get]
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Static("/docs", "./docs")
	url := ginSwagger.URL("http://localhost:8080/docs/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}
