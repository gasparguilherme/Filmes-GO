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

	dbURL := "postgres://postgres:senha@localhost:5433/films_go"

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		slog.Error("Erro ao conectar no banco", "error", err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	//testa a conexao
	if err := conn.Ping(ctx); err != nil {
		slog.Error("Error ao fazer ping no banco de dados", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("Conexão estabelcida com sucesso")

	userHandler := initUser(conn)
	filmHandler := initFilm(conn)

	api.StartApp(userHandler, filmHandler)

}
