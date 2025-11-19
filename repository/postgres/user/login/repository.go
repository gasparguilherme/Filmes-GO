package login

import "github.com/jackc/pgx/v5"

type Repository struct {
	connectionInstance *pgx.Conn
}

func NewPostgresRepository(conn *pgx.Conn) Repository {
	return Repository{
		connectionInstance: conn,
	}
}
