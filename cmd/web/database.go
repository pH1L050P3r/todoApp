package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(app *AppConfig) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: app.ENVIRONMENT.RDS_CONNECT_STRING,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect to database")
	}
	app.DB = db
}

func InitialMigrations(app *AppConfig) {
	//TODO:Add models here for migration
}
