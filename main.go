package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gasparguilherme/my-repository/api"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	userHandler := initUser()
	filmHandler := initFilm()

	dbURL := "postgres://postgres:senha@localhost:5433/films_go"

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		slog.Error("Erro ao conectar no banco", "error", err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	slog.Info("Conexão estabelcida com sucesso")

	api.StartApp(userHandler, filmHandler)

}
