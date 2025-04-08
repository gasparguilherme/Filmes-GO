package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func (Usecase) CreateFilm(userID int, title string, director string, year int, gender string) *entities.Film {
	newFilm := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Gender:   gender,
		UserID:   userID,
		ID:       film.GetNextFilmID(),
	}
	filmRepository := film.Repository{}
	filmRepository.SaveFilms(newFilm)

	return &newFilm

}
