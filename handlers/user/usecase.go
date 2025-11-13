package user

import "github.com/gasparguilherme/my-repository/domain/entities"

type Usecase interface {
	NewUser(name, email, password string) (*entities.User, error)
}
