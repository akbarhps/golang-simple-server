package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"time"
	"vps_server_playground/domain/mahasiswa"
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
	log.Println("Starting application...")

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

	mhsRepository := mahasiswa.NewRepository(db)
	mhsService := mahasiswa.NewService(mhsRepository)
	mhsController := mahasiswa.NewController(mhsService)

	app := fiber.New()

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	v1Group := app.Group("/api/v1")
	v1Group.Post("/mahasiswa", mhsController.CreateMahasiswa)

	log.Fatalln(app.Listen(config.ServerAddress))
}
