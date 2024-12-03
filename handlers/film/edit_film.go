package film

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func editFilm(userID int, filmID int, title string, director string, year int, gender string) {

	err := validate.ValidateFilm(title, director, year, gender)
	if err != nil {
		fmt.Println(err)
		return
	}

	film := usecases.EditFilm(userID, filmID, title, director, year, gender)
	fmt.Println("film successfully edited", film)

}
