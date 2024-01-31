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
	h.app.POST("/api/login", h.HandleUserLogin)
}

func (h LoginHandler) HandleLoginShow(c echo.Context) error {
	return render(c, login.LoginShow())
}

func (h LoginHandler) HandleUserLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	data := login.NewLoginFormData(login.WithUsername(username))

	if username == "" {
		return render(c, login.LoginForm(login.ErrorNoUsername, *data))
	}
	if password == "" {
		return render(c, login.LoginForm(login.ErrorNoPassword, *data))
	}

	return render(c, login.LoginForm(login.ErrorNone, *data))
}
