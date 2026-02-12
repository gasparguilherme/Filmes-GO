package user

import "github.com/gasparguilherme/my-repository/domain/entities"

type AuthResponse struct {
	User  entities.User `json:"user"`
	Token string        `json:"token"`
}
