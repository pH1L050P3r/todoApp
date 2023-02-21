package config

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/pH1L050P3r/todoApp/pkg/objects"
	"github.com/pH1L050P3r/todoApp/pkg/user"
	"github.com/spf13/viper"
)

func Init(Config *objects.AppConfig) {
	InitEnvironment(&Config.ENVIRONMENT)
	SetupServer(Config)
}

func InitEnvironment(e *objects.Environment) {
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

func SetRepository(app *objects.AppConfig) {
	user.SetRepository(app)
}

func SetupServer(app *objects.AppConfig) {
	app.SERVER = gin.Default()
}
