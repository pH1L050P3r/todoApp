package config

import (
	"log"

	"github.com/pH1L050P3r/todoApp/pkg/objects"
	"github.com/pH1L050P3r/todoApp/pkg/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(app *objects.AppConfig) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: app.ENVIRONMENT.RDS_CONNECT_STRING,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect to database")
	}
	app.DB = db
}

func InitialMigrations(app *objects.AppConfig) {
	app.DB.AutoMigrate(&user.User{})
}
