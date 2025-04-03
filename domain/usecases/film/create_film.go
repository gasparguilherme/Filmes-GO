package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/film"
)

func CreateFilm(userID int, title string, director string, year int, gender string) *entities.Film {
	film := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Gender:   gender,
		UserID:   userID,
		ID:       repository.GetNextFilmID(),
	}
	repository.SaveFilms(film)
	return &film

}
