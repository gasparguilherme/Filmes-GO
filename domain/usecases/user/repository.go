package user

import "github.com/gasparguilherme/my-repository/domain/entities"

type Repository interface {
	SaveUser(data entities.User) (int, error)
	DeleteUser(id int) error
}
