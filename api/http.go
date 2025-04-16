package api

import (
	"net/http"
)

func StartApp(userHandler UserHandler, filmHandler FilmHandler) {
	mux := http.NewServeMux()

	mux.Handle("POST /user", http.HandlerFunc(userHandler.CreateUser))
	mux.Handle("POST /film", http.HandlerFunc(filmHandler.HandleCreateFilm))
	mux.Handle("PUT /film", http.HandlerFunc(filmHandler.HandleEditFilm))

	mux.Handle("GET /film/{id}", http.HandlerFunc(filmHandler.HandleGetFilmByID))
	mux.Handle("GET /film", http.HandlerFunc(filmHandler.HandleGetFilmsByUserID))

	http.ListenAndServe(":8080", mux)
}
