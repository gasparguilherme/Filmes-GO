package main

import (
	"github.com/gasparguilherme/my-repository/api"
)

func main() {
	userHandler := initUser()
	filmHandler := initFilm()

	api.StartApp(userHandler, filmHandler)
}
