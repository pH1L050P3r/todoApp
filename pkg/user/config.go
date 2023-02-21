package user

import "github.com/pH1L050P3r/todoApp/pkg/objects"

var Repository *objects.Repository

func SetRepository(app *objects.AppConfig) {
	var repo = &objects.Repository{
		App: app,
	}
	Repository = repo
}
