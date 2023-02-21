package config

import (
	"github.com/pH1L050P3r/todoApp/pkg/objects"
	"github.com/pH1L050P3r/todoApp/pkg/user"
)

func InitRoutes(Config *objects.AppConfig) {
	Config.SERVER.POST("/user", user.CreateUser)
}
