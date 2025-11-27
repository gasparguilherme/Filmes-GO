package api

import (
	"net/http"
)

func StartApp(userHandler UserHandler, filmHandler FilmHandler, loginHandler LoginHandler) {
	mux := http.NewServeMux()

	mux.Handle("POST /user", http.HandlerFunc(userHandler.CreateUser))
	mux.Handle("POST /login", http.HandlerFunc(loginHandler.LoginHandler))
	mux.Handle("DELETE /user/{id}", http.HandlerFunc(userHandler.DeleteUser))

	mux.Handle("POST /film", http.HandlerFunc(filmHandler.CreateFilm))
	mux.Handle("PUT /edit-film", http.HandlerFunc(filmHandler.EditFilm))
	mux.Handle("GET /film/{id}", http.HandlerFunc(filmHandler.GetFilmByID))
	mux.Handle("GET /film", http.HandlerFunc(filmHandler.GetFilmsByUserID))
	mux.Handle("DELETE /film/{id}/user/{userID}", http.HandlerFunc(filmHandler.DeleteFilm))

	http.ListenAndServe(":8080", mux)
}
