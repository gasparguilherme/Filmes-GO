package user

import "github.com/gasparguilherme/my-repository/domain/entities"

type Repository interface {
	GetNextUserID() int
	SaveUser(data entities.User)
}
