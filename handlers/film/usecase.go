package film

import "github.com/gasparguilherme/my-repository/domain/entities"

type Usecase interface {
	EditFilm(filmID int, newTitle, newDirector string, newYear int, newGenre string, userID int) error
	FindFilmsByUserID(id int) ([]entities.Film, error)
	FindFilmByID(id int) (*entities.Film, error)
	CreateFilm(userID int, title string, director string, year int, genre string) (*entities.Film, error)
}
