package login

import "github.com/gasparguilherme/my-repository/domain/entities"

type Usecase interface {
	Login(data entities.Login) (*entities.User, error)
}
