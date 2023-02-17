package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Environment struct {
	ENVIRONMENT string

	WEB_HOST string
	WEB_PORT string

	RDS_USERNAME       string
	RDS_PASSWORD       string
	RDS_HOSTNAME       string
	RDS_PORT           string
	RDS_DB_NAME        string
	RDS_CONNECT_STRING string
}

type AppConfig struct {
	ENVIRONMENT Environment
	DB          *gorm.DB
}

func (Config *AppConfig) Init() {
	InitEnvironment(&Config.ENVIRONMENT)
}

func InitEnvironment(e *Environment) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(".env not found in environment")
	}

	value, ok := viper.Get("ENVIRONMENT").(string)
	if !ok {
		log.Printf("DEBUG : ENVIRONMENT variable not found")
		log.Printf("DEBUG : ENVIRONMENT set to test")
		e.ENVIRONMENT = "test"
	} else {
		e.ENVIRONMENT = value
	}

	var database_variables = []string{
		"RDS_USERNAME",
		"RDS_PASSWORD",
		"RDS_HOSTNAME",
		"RDS_PORT",
		"RDS_DB_NAME",
	}
	for _, variable := range database_variables {
		value, ok = viper.Get(variable).(string)
		if !ok {
			log.Fatal(fmt.Printf("%s is not found", variable))
		} else {
			reflect.ValueOf(e).Elem().FieldByName(variable).SetString(value)
			log.Printf("%s : %s", variable, value)
		}
	}

	e.RDS_CONNECT_STRING = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		e.RDS_HOSTNAME,
		e.RDS_USERNAME,
		e.RDS_PASSWORD,
		e.RDS_DB_NAME,
		e.RDS_PORT,
	)

}
