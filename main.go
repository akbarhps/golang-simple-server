package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	return db, nil
}

func main() {
	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("error read config: ", err)
	}

	db, err := getConnection(&config)
	if err != nil {
		log.Fatal("error get connection: ", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("error get sql instance: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(config.DBMigrationSource, config.DBDriver, driver)
	if err != nil {
		log.Fatal("error create migrate instance: ", err)
	}

	if err := m.Up(); err != nil {
		log.Println("error migrating: ", err)
	}
}
