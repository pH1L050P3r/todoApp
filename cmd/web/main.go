package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var AppSettings AppConfig
	AppSettings.Init()

	log.Println(AppSettings.ENVIRONMENT)

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	serverURL := "0.0.0.0:8000"
	server.Run(serverURL)
}
