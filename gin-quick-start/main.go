package main

import (
	// "net/http"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}			

func main() {

	router := gin.Default()
	router.GET("/ping", pong)
	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}