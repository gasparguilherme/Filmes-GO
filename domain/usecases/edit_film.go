package usecases

import (
	"github.com/gasparguilherme/my-repository/repository/film"
)

func EditFilm(filmID int, newTitle, newDirector string, newYear int, newGender string) error {
	err := film.UpdateFilm(filmID, newTitle, newDirector, newYear, newGender)
	return err
}
