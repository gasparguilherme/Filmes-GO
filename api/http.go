package api

import (
	"net/http"
)

func StartApp(userHandler UserHandler, filmHandler FilmHandler) {
	mux := http.NewServeMux()

	mux.Handle("POST /user", http.HandlerFunc(userHandler.CreateUser))
	mux.Handle("POST /film", http.HandlerFunc(filmHandler.CreateFilm))
	mux.Handle("PUT /film", http.HandlerFunc(filmHandler.EditFilm))

	mux.Handle("GET /film/{id}", http.HandlerFunc(filmHandler.GetFilmByID))
	mux.Handle("GET /film", http.HandlerFunc(filmHandler.GetFilmsByUserID))

	http.ListenAndServe(":8080", mux)
}
