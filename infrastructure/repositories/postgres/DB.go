package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnection() (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		"postgresql://postgres:password@localhost:5432/ddcomments?sslmode=disable",
	)
}
