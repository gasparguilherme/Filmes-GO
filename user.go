package main

import (
	usecase "github.com/gasparguilherme/my-repository/domain/usecases/user"
	handler "github.com/gasparguilherme/my-repository/handlers/user"
	repository "github.com/gasparguilherme/my-repository/repository/user"
)

func initUser() handler.Handler {
	r := repository.NewRepository()
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)
	return h

}
