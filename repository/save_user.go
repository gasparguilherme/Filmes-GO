package repository

import "github.com/gasparguilherme/my-repository/domain/entities"

var listaUsuarios []entities.Usuario

func Save_user(dados entities.Usuario) {
	listaUsuarios = append(listaUsuarios, dados)
}
