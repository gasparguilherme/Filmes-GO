package api

import (
	"net/http"
)

type FilmHandler interface {
	EditFilm(w http.ResponseWriter, r *http.Request)
	GetFilmByID(w http.ResponseWriter, r *http.Request)
	GetFilmsByUserID(w http.ResponseWriter, r *http.Request)
	CreateFilm(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
