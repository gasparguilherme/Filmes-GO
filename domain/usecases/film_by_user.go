package usecases

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/user"
)

func FilmByUser(id int) ([]entities.Film, error) {
	films, err := user.FindFilmsByUserID(id)
	if err != nil {
		return nil, err
	}
	return films, nil
}
