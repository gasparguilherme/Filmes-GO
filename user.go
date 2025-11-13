package main

import (
	usecase "github.com/gasparguilherme/my-repository/domain/usecases/user"
	handler "github.com/gasparguilherme/my-repository/handlers/user"
	"github.com/gasparguilherme/my-repository/repository/postgres/user"
	"github.com/jackc/pgx/v5"
)

func InitUser(conn *pgx.Conn) handler.Handler {
	r := user.NewPostgresRepository(conn)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)
	return h

}
