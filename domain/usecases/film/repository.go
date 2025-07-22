package film

import "github.com/gasparguilherme/my-repository/domain/entities"

type Repository interface {
	UpdateFilm(filmID int, newTitle string, newDirector string, newYear int, newGender string) error
	FindFilmsByUserID(userID int) ([]entities.Film, error)
	GetNextFilmID() int
	SaveFilm(data entities.Film) (int, error)
	FindFilmByID(id int) (entities.Film, error)
}
