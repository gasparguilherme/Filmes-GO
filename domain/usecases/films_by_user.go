package usecases

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func FilmsByUser(id int) ([]entities.Film, error) {
	for _, user := range repository.GetUsers() {
		if user.ID == id {
			return user.Films, nil
		}
	}
	return nil, errors.New("user not found")
}
