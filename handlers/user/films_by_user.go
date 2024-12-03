package user

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func filmByUSer(id int) {
	err := validate.ValidateID(id)
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
