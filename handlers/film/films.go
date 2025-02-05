package film

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func CreateFilm(jsonInput []byte) {
	var film *entities.Film
	err := json.Unmarshal(jsonInput, &film)
	if err != nil {
		fmt.Println("error unmarshal json")
		return
	}

	err = validate.ValidateFilm(film.Title, film.Director, film.Year, film.Gender)
	if err != nil {
		fmt.Println("error validating filme", err)
		return
	}

	film = usecases.CreateFilm(film.UserID, film.Title, film.Director, film.Year, film.Gender)

	jsonFilm, err := json.Marshal(film)
	if err != nil {
		fmt.Println("error converting json format ", err)
		return
	}

	fmt.Println("film successfully created!", string(jsonFilm))

}
