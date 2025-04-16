package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) FindFilmByID(id int) (entities.Film, error) {
	filmFound, err := u.repository.FindFilmByID(id)
	if err != nil {
		return entities.Film{}, err
	}
	return filmFound, nil
}
