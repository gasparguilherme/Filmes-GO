package film

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/user"
)

func (Repository) FindFilmsByUserID(id int) ([]entities.Film, error) {

	for _, user := range user.GetUsers() {
		if user.ID == id {
			for _, film := range GetFilms() {
				if film.UserID == id {
				}
			}
			return GetFilms(), nil
		}
	}
	return nil, errors.New("user not found")
}
