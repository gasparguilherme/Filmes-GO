package film

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/film"
)

func (Usecase) CreateFilm(userID int, title string, director string, year int, gender string) *entities.Film {
	film := entities.Film{
		Title:    title,
		Director: director,
		Year:     year,
		Gender:   gender,
		UserID:   userID,
		ID:       repository.GetNextFilmID(),
	}
	filmRepository := repository.Repository{}
	filmRepository.SaveFilms(film)

	return &film

}
