package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// Inicjuję serwer
	router := gin.Default()

	// Definiuję prosty route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Maykie",
		})
	})

	// Odpalam serwer
	router.Run(":8080")
}
