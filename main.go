package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Get the port from command line arguments
	port := "8080" // default port
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	r.Run(":" + port)
}
