package film

import "github.com/gasparguilherme/my-repository/domain/entities"

type Repository interface {
	UpdateFilm(data entities.Film) error
	FindFilmsByUserID(userID int) ([]entities.Film, error)
	SaveFilm(data entities.Film) (int, error)        // corrigido
	FindFilmByID(filmID int) (*entities.Film, error) // corrigido
}
