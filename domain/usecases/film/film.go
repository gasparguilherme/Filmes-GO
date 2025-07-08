package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) CreateFilm(userID int, title string, director string, year int, genre string) *entities.Film {
	newFilm := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Genre:    genre,
		UserID:   userID,
	}
	u.repository.SaveFilms(newFilm)

	return &newFilm

}
