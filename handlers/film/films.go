package film

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func createFilm(title string, director string, year int, gender string) {

	err := validate.ValidateFilm(title, director, year, gender)
	if err != nil {
		fmt.Println(err)
	}

	film := usecases.CreateFilm(title, director, year, gender)
	fmt.Println("film successfully created!", film)

}
