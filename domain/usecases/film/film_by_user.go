package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func (Usecase) FilmByUser(id int) ([]entities.Film, error) {
	films, err := film.FindFilmsByUserID(id)
	if err != nil {
		return nil, err
	}
	return films, nil
}
