package login

import "github.com/gasparguilherme/my-repository/domain/entities"

type Repository interface {
	FindUserByEmail(email string) (*entities.User, error)
}
