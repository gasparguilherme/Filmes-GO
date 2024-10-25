package usecases

import (
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/repository"
)

func criarFilme(titulo string, diretor string, ano int, genero string) entities.Filme {
	filme := entities.Filme{
		Titulo:  titulo,
		Diretor: diretor,
		Ano:     ano,
		Genero:  genero,
		ID:      repository.GetNextID_films(),
	}
	return filme

}
