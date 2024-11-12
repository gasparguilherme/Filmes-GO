package usecases

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func FilmsByID(id int) (entities.Film, *entities.User, error) {
	for _, user := range repository.GetUsers() {
		for _, film := range user.Films {
			if film.ID == id {
				return film, &entities.User{}, nil
			}
		}
	}
	return entities.Film{}, nil, errors.New("user not found")
}
