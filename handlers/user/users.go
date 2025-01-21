package user

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func createUser(jsonInput []byte) {
	var user entities.User
	err := json.Unmarshal(jsonInput, &user)
	if err != nil {
		fmt.Println("error unmarshal json", err)
		return
	}

	err = validate.ValidateUser(user.Name, user.Email)
	if err != nil {
		fmt.Println("validation error", err)
		return
	}

	user = *usecases.NewUser(user.Email, user.Name)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error converting json format", err)
		return
	}

	fmt.Println("Json generated:", string(jsonUser))

}
