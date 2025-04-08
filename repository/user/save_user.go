package user

import "github.com/gasparguilherme/my-repository/domain/entities"

var listUsers []entities.User

func (Repository) SaveUser(data entities.User) {
	listUsers = append(listUsers, data)
}

func GetUsers() []entities.User {
	return listUsers
}
