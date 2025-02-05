package film

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func FilmByID(jsonInput []byte) {
	var f entities.Film
	err := json.Unmarshal(jsonInput, &f)
	if err != nil {
		fmt.Println("error unmarshal json", err)
	}

	err = validate.ValidateID(f.ID)
	if err != nil {
		fmt.Println("validation error", err)
		return
	}

	f, err = film.FilmByID(f.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonFilm, err := json.Marshal(f)
	if err != nil {
		fmt.Println("error converting json format ", err)
		return
	}

	fmt.Println("Film found", string(jsonFilm))
}
