package main

import (
	usecase "github.com/gasparguilherme/my-repository/domain/usecases/user/login"
	handler "github.com/gasparguilherme/my-repository/handlers/user/login"
	repository "github.com/gasparguilherme/my-repository/repository/postgres/user/login"
	"github.com/jackc/pgx/v5"
)

func InitLogin(conn *pgx.Conn) handler.Handler {
	r := repository.NewPostgresRepository(conn)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)
	return h
}
