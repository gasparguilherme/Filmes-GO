package user

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) NewUser(name, email string) *entities.User {
	newUser := entities.User{
		Name:  name,
		Email: email,
		ID:    u.repository.GetNextUserID(),
	}
	u.repository.SaveUser(newUser)

	return &newUser

}
