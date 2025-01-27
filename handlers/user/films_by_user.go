package user

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func filmByUSer(jsonInput []byte) {
	var User struct {
		UserID int `json:"user_id"`
	}

	err := json.Unmarshal(jsonInput, &User)
	if err != nil {
		fmt.Println("error unmarshal json", err)
	}

	err = validate.ValidateID(User.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := usecases.FilmByUser(User.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user id found", user)
}
