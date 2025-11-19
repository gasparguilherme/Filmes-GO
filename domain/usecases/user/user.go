package user

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

func (u Usecase) NewUser(name, email, password string) (*entities.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := entities.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	id, err := u.repository.SaveUser(newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = id
	return &newUser, nil

}
