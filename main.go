package main

import (
	"net/http"

	"github.com/gasparguilherme/my-repository/handlers/film"
	"github.com/gasparguilherme/my-repository/handlers/user"
)

func main() {
	http.Handle("/users/create-user", http.HandlerFunc(user.CreateUser))

	http.Handle("/films/create-film", http.HandlerFunc(film.CreateFilm))

	http.Handle("/films/edit-film", http.HandlerFunc(film.HandleEditFilm))

	http.Handle("/films/by-id", http.HandlerFunc(film.HandleGetFilmByID))

	http.Handle("/users/by-user", http.HandlerFunc(user.HandleGetFilmsByUserID))

	http.ListenAndServe(":8090", nil)

}
