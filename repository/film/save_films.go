package film

import "github.com/gasparguilherme/my-repository/domain/entities"

var listFilms []entities.Film

func SaveFilms(data entities.Film) {
	listFilms = append(listFilms, data)
}

func GetFilms() []entities.Film {
	return listFilms
}
