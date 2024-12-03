package user

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func createUser(name, email string) {

	err := validate.ValidateUser(name, email)
	if err != nil {
		fmt.Println(err)
		return
	}

	user := usecases.NewUser(name, email)
	fmt.Println("user created successfully", user)

}
