package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gasparguilherme/my-repository/api"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Erro ao carregar .env", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()

	//pego o ambiente
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		slog.Error("DB_URL não encontrada no .env")
		os.Exit(1)
	}

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

	userHandler := InitUser(conn)
	filmHandler := InitFilm(conn)
	loginHandler := InitLogin(conn)

	api.StartApp(userHandler, filmHandler, loginHandler)

}
