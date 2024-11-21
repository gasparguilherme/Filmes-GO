package usecases

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

type userLocal entities.User

func (u *userLocal) editFilm(filmID int, tittle string, director string, year int, gender string) error {
	for i, filme := range u.Films {
		if filme.ID == filmID {
			u.Films[i].Title = tittle
			u.Films[i].Director = director
			u.Films[i].Year = year
			u.Films[i].Gender = gender
			return nil

		}
	}
	return fmt.Errorf("film with ID not found")
}
