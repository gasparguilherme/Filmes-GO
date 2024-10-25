package usecases

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository"
)

func NovoUsuario(nome, email string) *entities.Usuario {
	usuario := entities.Usuario{
		Nome:  nome,
		Email: email,
		ID:    repository.GetNextID_users(),
	}
	repository.Save_user(usuario)

	return &usuario

}
