package login

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases/user/login"
)

type Usecase interface {
	Login(data login.Login) (*entities.User, error)
}
