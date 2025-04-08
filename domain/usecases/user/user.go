package user

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository/user"
)

func (Usecase) NewUser(name, email string) *entities.User {
	newUser := entities.User{
		Name:  name,
		Email: email,
		ID:    user.GetNextUserID(),
	}
	userRepository := user.Repository{}
	userRepository.SaveUser(newUser)

	return &newUser

}
