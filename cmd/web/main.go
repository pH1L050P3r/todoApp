package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pH1L050P3r/todoApp/pkg/config"
	"github.com/pH1L050P3r/todoApp/pkg/objects"
)

func main() {
	var AppSettings objects.AppConfig

	config.Init(&AppSettings)
	config.ConnectDB(&AppSettings)
	config.InitialMigrations(&AppSettings)
	config.SetRepository(&AppSettings)
	config.InitRoutes(&AppSettings)

	log.Println("Environment :", AppSettings.ENVIRONMENT)

	AppSettings.SERVER.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	serverURL := "0.0.0.0:8000"
	AppSettings.SERVER.Run(serverURL)
}
