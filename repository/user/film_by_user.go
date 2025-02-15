package user

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func FindFilmsByUserID(id int) ([]entities.Film, error) {

	for _, user := range GetUsers() {
		if user.ID == id {
			for _, film := range film.GetFilms() {
				if film.UserID == id {
				}
			}
			return film.GetFilms(), nil
		}
	}
	return nil, errors.New("user not found")
}
