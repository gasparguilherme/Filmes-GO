package main

import (
	usecase "github.com/gasparguilherme/my-repository/domain/usecases/film"
	handler "github.com/gasparguilherme/my-repository/handlers/film"
	repository "github.com/gasparguilherme/my-repository/repository/film"
)

func initFilm() handler.Handler {
	r := repository.NewRepository()
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)

	return h

}
