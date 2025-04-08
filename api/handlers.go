package api

import "github.com/gasparguilherme/my-repository/domain/entities"

type FilmHandler interface {
	EditFilm(filmID int, newTitle, newDirector string, newYear int, newGender string)
	FilmByID(id int) (entities.Film, error)
	FilmByUser(id int) ([]entities.Film, error)
	CreateFilm(userID int, title string, director string, year int, gender string)
}

type UserHandler interface {
	NewUser(name, email string) *entities.User
}
