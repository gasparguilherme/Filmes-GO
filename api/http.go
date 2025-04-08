package api

import (
	"net/http"

	"github.com/gasparguilherme/my-repository/handlers/film"
	"github.com/gasparguilherme/my-repository/handlers/user"
)

func StartApp() {
	mux := http.NewServeMux()
	userHandler := user.Handler{}
	filmHandler := film.Handler{}

	mux.Handle("POST /user", http.HandlerFunc(userHandler.CreateUser))
	mux.Handle("POST /film", http.HandlerFunc(filmHandler.CreateFilm))
	mux.Handle("PUT /film", http.HandlerFunc(filmHandler.HandleEditFilm))

	mux.Handle("GET /film/{id}", http.HandlerFunc(filmHandler.HandleGetFilmByID))
	mux.Handle("GET /film", http.HandlerFunc(filmHandler.HandleGetFilmsByUserID))

	http.ListenAndServe(":8080", mux)
}
