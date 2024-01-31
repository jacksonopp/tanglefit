package main

import (
	"net/http"

	"github.com/jacksonopp/tanglefit/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Debug = true

	app.Static("/static", "static")

	loginHandler := handlers.NewLoginHandler(app)
	loginHandler.HandleAllRoutes()
	// app.GET("/login", loginHandler.HandleLoginShow)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	app.Start(":3000")
}
