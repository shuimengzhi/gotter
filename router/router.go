package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get pong",
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post pong",
		})
	})
	return r
}
