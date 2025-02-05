package main

import (
	"github.com/gasparguilherme/my-repository/handlers/film"
	"github.com/gasparguilherme/my-repository/handlers/user"
)

func main() {

	//userJson variavel do tipo []byte guarda o json gerado e passado como argumento na funcao CreateUSer da cmd handlers
	var userJson []byte = []byte(`{
		"name": "Guilherme",
		"email": "guilherme474@gmail.com"
		}`)
	user.CreateUser(userJson)

	var filmJson []byte = []byte(`{
		"user_id":1,
		"title": "Star Wars",
		"director": "George Lucas",
		"year": 1970,
		"gender": "Ficção Cientifica"
		}`)
	film.CreateFilm(filmJson)

	var filmByUserJson []byte = []byte(`{
	"user_id": 1}`)
	user.FilmByUser(filmByUserJson)

	var filmByIDJson []byte = []byte(`{
	"id": 1}`)
	film.FilmByID(filmByIDJson)

	var editFilmJson []byte = []byte(`{
	"user_id":1,
	"title": "Star Wars",
	"director": "George Lucas",
	"year": 1977,
	"gender": "Ficção Cientifica"
		}`)
	film.EditFilm(editFilmJson)
}
