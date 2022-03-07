package postgres

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"warehouse/shared"
)

type PostgresBase struct {
	DB *gorm.DB
}

func NewPostgresBase() *PostgresBase {
	connectionDb, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	return &PostgresBase{connectionDb}
}

func getConnection() (*gorm.DB, error) {
	uri := shared.GetEnv("DATABASE_URI")
	return gorm.Open(postgres.Open(uri), &gorm.Config{})
}