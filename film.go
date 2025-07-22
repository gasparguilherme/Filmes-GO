package main

import (
	usecase "github.com/gasparguilherme/my-repository/domain/usecases/film"
	handler "github.com/gasparguilherme/my-repository/handlers/film"
	repository "github.com/gasparguilherme/my-repository/repository/postgres/film"
	"github.com/jackc/pgx/v5"
)

func initFilm(conn *pgx.Conn) handler.Handler {
	r := repository.NewPostgresRepository(conn)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)
	return h

}
