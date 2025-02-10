package usecases

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func FilmByID(id int) (entities.Film, error) {

	for _, films := range film.GetFilms() {
		if films.ID == id {
			return films, nil
		}
	}

	return entities.Film{}, errors.New("filme nao encontrado")
}
