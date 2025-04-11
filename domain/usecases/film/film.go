package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) CreateFilm(userID int, title string, director string, year int, gender string) *entities.Film {
	newFilm := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Gender:   gender,
		UserID:   userID,
		ID:       u.repository.GetNextFilmID(),
	}
	u.repository.SaveFilms(newFilm)

	return &newFilm

}
