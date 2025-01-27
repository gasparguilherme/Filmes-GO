package film

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func filmByID(jsonInput []byte) {
	var film entities.Film
	err := json.Unmarshal(jsonInput, &film)
	if err != nil {
		fmt.Println("error unmarshal json", err)
	}

	err = validate.ValidateID(film.ID)
	if err != nil {
		fmt.Println("validation error", err)
		return
	}

	film, err = usecases.FilmByID(film.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonFilm, err := json.Marshal(film)
	if err != nil {
		fmt.Println("error converting json format ", err)
		return
	}

	fmt.Println("Film found", string(jsonFilm))
}
