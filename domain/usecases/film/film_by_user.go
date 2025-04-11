package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) FilmByUser(id int) ([]entities.Film, error) {
	filmFound, err := u.repository.FindFilmsByUserID(id)
	if err != nil {
		return nil, err
	}
	return filmFound, nil

}
