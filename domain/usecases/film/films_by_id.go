package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func FilmByID(id int) (entities.Film, error) {
	film, err := film.FindFilmByID(id)
	if err != nil {
		return entities.Film{}, err
	}
	return film, nil
}
