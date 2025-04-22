package api

import (
	"net/http"
)

type FilmHandler interface {
	HandleEditFilm(w http.ResponseWriter, r *http.Request)
	HandleGetFilmByID(w http.ResponseWriter, r *http.Request)
	HandleGetFilmsByUserID(w http.ResponseWriter, r *http.Request)
	HandleCreateFilm(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
