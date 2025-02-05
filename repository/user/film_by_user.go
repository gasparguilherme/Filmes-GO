package user

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func FilmByUser(id int) ([]entities.Film, error) {
	var userFilms []entities.Film

	for _, user := range GetUsers() {
		if user.ID == id {
			for _, film := range film.GetFilms() {
				if film.UserID == id {
					userFilms = append(userFilms, film)
				}
			}
			return userFilms, nil
		}
	}
	return nil, errors.New("user not found")
}
