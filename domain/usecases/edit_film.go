package usecases

import (
	"fmt"

	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func EditFilm(userID int, filmID int, title string, director string, year int, gender string) error {
	users := repository.GetUsers()

	for _, user := range users {
		if user.ID == userID {
			for i, film := range user.Films {
				if film.ID == filmID {
					user.Films[i].Title = title
					user.Films[i].Director = director
					user.Films[i].Year = year
					user.Films[i].Gender = gender

					return nil
				}
			}
		}
	}
	return fmt.Errorf("film with ID not found")

}
