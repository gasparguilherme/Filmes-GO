package user

import "github.com/jackc/pgx/v5"

type Repository struct {
	conn pgx.Conn
}
