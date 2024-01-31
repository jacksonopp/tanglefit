package handlers

import (
	"github.com/jacksonopp/tanglefit/view/login"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	app *echo.Echo
}

func NewLoginHandler(app *echo.Echo) *LoginHandler {
	return &LoginHandler{app}
}

func (h LoginHandler) HandleAllRoutes() {
	h.app.GET("/login", h.HandleLoginShow)
}

func (h LoginHandler) HandleLoginShow(c echo.Context) error {
	return render(c, login.LoginShow())
}
