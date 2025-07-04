package user

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (u Usecase) NewUser(name, email string) (*entities.User, error) {
	newUser := entities.User{
		Name:  name,
		Email: email,
	}
	id, err := u.repository.SaveUser(newUser)
	if err != nil {
		return nil, err
	}
	newUser.ID = id
	return &newUser, nil

}
