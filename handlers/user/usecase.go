package user

import "github.com/gasparguilherme/my-repository/domain/entities"

type Usecase interface {
	NewUser(name, email string) (*entities.User, error)
}
