package usecases

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/film"
)

func CreateFilm(title string, director string, year int, gender string) *entities.Film {
	film := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Gender:   gender,
		ID:       repository.GetNextFilmID(),
	}
	return &film

}
