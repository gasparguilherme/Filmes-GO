package film

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func filmByID(id int) {
	err := validate.ValidateID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	film, err := usecases.FilmByID(id)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Film found", film)
}
