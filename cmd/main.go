package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jacksonopp/tanglefit/db"
	"github.com/jacksonopp/tanglefit/handlers"
	"github.com/labstack/echo/v4"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	_, err := migrateDb()
	if err != nil && err != migrate.ErrNoChange {
		log.Error("error migrating database:", err)
		os.Exit(1)
	}

	sqldb, err := sql.Open("postgres", "port=5438 user=postgres password=postgres dbname=tanglefit sslmode=disable")
	if err != nil {
		log.Error("error setting up database", err)
		os.Exit(1)
	}
	defer sqldb.Close()

	queries := db.New(sqldb)
	ctx := context.Background()

	app := setupApp()
	if l, ok := app.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
		l.SetLevel(log.DEBUG)
	}

	handlers.NewLoginHandler(app, queries, ctx).HandleAllRoutes()
	handlers.NewSignupHandler(app, queries, ctx).HandleAllRoutes()

	app.Start(":3000")
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
