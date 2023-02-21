package objects

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppConfig struct {
	ENVIRONMENT Environment
	DB          *gorm.DB
	SERVER      *gin.Engine
}
