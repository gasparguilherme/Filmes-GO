package handlers

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
)

func filmByID(id int) {
	err := validateID(id)
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
