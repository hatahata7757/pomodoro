package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/greeting", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		c.JSON(200, gin.H{
			"message": "hello, world",
		})
	})
	r.Run(":8080")
}
