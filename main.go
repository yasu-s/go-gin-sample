package main

import (
	"github.com/gin-gonic/gin"
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
	r.Run()
}
