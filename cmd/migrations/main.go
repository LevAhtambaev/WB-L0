package main

import (
	"LO/pkg/dsn"
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"
)

const (
	migrationsPath = "migrations"
	driver         = "postgres"
)

func main() {

	ctx := context.Background()
	log.Info("Starting migrations")

	err := godotenv.Load()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("No .env file loaded")
	}

	dbDSN := dsn.FromEnv()

	db, err := sql.Open(driver, dbDSN)
	if err != nil {
		log.Errorf("Failed to connect to db: %s", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	log.Info("The database connection was established successfully")

	_ = goose.SetDialect(driver)

	log.Info("Upping migrations")
	err = goose.Up(db, migrationsPath)
	if err != nil {
		log.Errorf("Failed to migrate: %v", err)
	}

	log.Info("DB migrations completed")
}
