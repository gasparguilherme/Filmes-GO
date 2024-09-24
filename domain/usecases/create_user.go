package usecases

import "github.com/gasparguilherme/my-repository/domain/entities"

func NovoUsuario(nome, email string) *entities.Usuario {
	usuario := &Usuario{
		Nome: nome,
		Email: email,
		ID: idAtual,
	}
	idAtual++
	listaUsuarios = append(listaUsuarios,usuario )
	
	return usuario

}