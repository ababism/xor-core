package middleware

import "github.com/gin-gonic/gin"

func Ping(g *gin.RouterGroup) {
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
