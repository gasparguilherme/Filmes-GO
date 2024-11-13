package usecases

import (
	"errors"

	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func FilmByID(id int) (entities.Film, error) {
	for _, user := range repository.GetUsers() {
		for _, film := range user.Films {
			if film.ID == id {
				return film, nil
			}
		}
	}
	return entities.Film{}, errors.New("film not found")
}
