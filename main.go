package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"log"
	"time"
)

func getConnection(config *Config) (*sql.DB, error) {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(30)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(config.DBMigrationSource, config.DBDriver, driver)
	if err != nil {
		return nil, err
	}

	if err := m.Steps(2); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("error read config:", err)
	}

	_, err = getConnection(&config)
	if err != nil {
		log.Fatal("error get connection:", err)
	}
}
