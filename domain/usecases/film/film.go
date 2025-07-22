package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) CreateFilm(userID int, title string, director string, year int, genre string) (*entities.Film, error) {
	newFilm := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Genre:    genre,
		User_ID:  userID,
	}

	id, err := u.repository.SaveFilm(newFilm)
	if err != nil {
		return nil, err
	}

	newFilm.ID = id
	return &newFilm, nil
}
