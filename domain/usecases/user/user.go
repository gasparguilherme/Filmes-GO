package user

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func (Usecase) NewUser(name, email string) *entities.User {
	user := entities.User{
		Name:  name,
		Email: email,
		ID:    repository.GetNextUserID(),
	}
	userRepository := repository.Repository{}
	userRepository.SaveUser(user)

	return &user

}
