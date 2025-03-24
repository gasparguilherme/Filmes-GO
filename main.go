package main

import (
	"net/http"

	"github.com/gasparguilherme/my-repository/handlers/film"
	"github.com/gasparguilherme/my-repository/handlers/user"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/users/create-user", http.HandlerFunc(user.CreateUser))
	mux.Handle("/films/create-film", http.HandlerFunc(film.CreateFilm))
	mux.Handle("/films/edit-film", http.HandlerFunc(film.HandleEditFilm))
	mux.Handle("/films/by-id", http.HandlerFunc(film.HandleGetFilmByID))
	mux.Handle("/users/by-user", http.HandlerFunc(user.HandleGetFilmsByUserID))

	http.ListenAndServe(":9990", mux)

}
