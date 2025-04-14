package film

import "github.com/gasparguilherme/my-repository/domain/entities"

type Usecase interface {
	UpdateFilm(filmID int, newTitle string, newDirector string, newYear int, newGender string) error
	FindFilmsByUserID(id int) ([]entities.Film, error)
	GetNextFilmID() int
	SaveFilms(data entities.Film)
	FindFilmByID(id int) (entities.Film, error)
	CreateFilm(userID int, title string, director string, year int, gender string) *entities.Film
}
