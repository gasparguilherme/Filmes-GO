package handlers

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
)

func filmByUSer(id int) {
	err := validateID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := usecases.FilmByUser(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user id found", user)
}
