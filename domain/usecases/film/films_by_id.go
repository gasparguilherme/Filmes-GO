package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func (Usecase) FilmByID(id int) (entities.Film, error) {
	filmRepository := film.Repository{}
	filmFound, err := filmRepository.FindFilmByID(id)
	if err != nil {
		return entities.Film{}, err
	}
	return filmFound, nil
}
