package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func (Usecase) FilmByUser(id int) ([]entities.Film, error) {
	filmRepository := film.Repository{}
	filmFound, err := filmRepository.FindFilmsByUserID(id)
	if err != nil {
		return nil, err
	}
	return filmFound, nil

}
