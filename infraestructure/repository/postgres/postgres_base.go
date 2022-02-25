package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"warehouse/shared"
)

type PostgresBase struct {
	DB *sql.DB
}

func NewPostgresBase() *PostgresBase {
	connectionDb, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	return &PostgresBase{connectionDb}
}

func getConnection() (*sql.DB, error) {
	uri := shared.GetEnv("DATABASE_URI")
	return sql.Open("postgres", uri)
}