package user

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func NewUser(name, email string) *entities.User {
	user := entities.User{
		Name:  name,
		Email: email,
		ID:    repository.GetNextUserID(),
	}
	repository.SaveUser(user)
	return &user

}
