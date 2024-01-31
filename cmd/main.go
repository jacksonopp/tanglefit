package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jacksonopp/tanglefit/db"
	"github.com/jacksonopp/tanglefit/handlers"
	"github.com/labstack/echo/v4"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// go:embed schema.sql
var ddl string

func main() {
	database, err := migrateDb()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln("error migrating database:", err)
	}

	queries, err := setupDb(database)
	if err != nil {
		log.Fatalln("error setting up database", err)
	}
	ctx := context.Background()

	app := setupApp()

	loginHandler := handlers.NewLoginHandler(app, queries, ctx)
	loginHandler.HandleAllRoutes()

	app.Start(":3000")
}

func setupDb(database *sql.DB) (*db.Queries, error) {
	sqldb, err := sql.Open("postgres", "port=5438 user=postgres password=postgres dbmane=tanglefit sslmode=disable")
	if err != nil {
		return nil, err
	}

	queries := db.New(sqldb)
	return queries, nil
}

func setupApp() *echo.Echo {
	app := echo.New()
	app.Debug = true
	app.Static("/static", "static")
	return app
}

func migrateDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5438/tanglefit?sslmode=disable")
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil {
		return nil, err
	}
	return db, nil
}
