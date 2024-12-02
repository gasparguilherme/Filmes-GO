package handlers

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
)

func createUser(name, email string) {

	err := validateUser(name, email)
	if err != nil {
		fmt.Println(err)
		return
	}

	user := usecases.NewUser(name, email)
	fmt.Println("user created successfully", user)

}
