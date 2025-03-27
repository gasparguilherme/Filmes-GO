package main

import (
	"net/http"

	"github.com/gasparguilherme/my-repository/handlers/film"
	"github.com/gasparguilherme/my-repository/handlers/user"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("POST /user", http.HandlerFunc(user.CreateUser))
	mux.Handle("POST /film", http.HandlerFunc(film.CreateFilm))
	mux.Handle("PUT /film", http.HandlerFunc(film.HandleEditFilm))

	mux.Handle("GET /film/{id}", http.HandlerFunc(film.HandleGetFilmByID))
	mux.Handle("GET /film", http.HandlerFunc(user.HandleGetFilmsByUserID))

	http.ListenAndServe(":8080", mux)

}
