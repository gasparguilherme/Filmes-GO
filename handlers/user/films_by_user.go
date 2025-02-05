package user

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/handlers/validate"
	"github.com/gasparguilherme/my-repository/repository/user"
)

type FilmByUserRequest struct {
	UserID int `json:"user_id"`
}

func FilmByUser(jsonInput []byte) {
	var User FilmByUserRequest

	err := json.Unmarshal(jsonInput, &User)
	if err != nil {
		fmt.Println("error unmarshal json", err)
	}

	err = validate.ValidateID(User.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	u, err := user.FilmByUser(User.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user id found", u)
}
