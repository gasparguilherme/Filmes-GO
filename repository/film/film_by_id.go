package film

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func FindFilmByID(id int) (entities.Film, error) {

	for _, films := range GetFilms() {
		if films.ID == id {
			return films, nil
		}
	}

	return entities.Film{}, errors.New("filme nao encontrado")
}
