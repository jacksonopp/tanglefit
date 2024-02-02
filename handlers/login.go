package handlers

import (
	"context"

	"github.com/jacksonopp/tanglefit/db"
	"github.com/jacksonopp/tanglefit/view/login"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	app *echo.Echo
	db  *db.Queries
	ctx context.Context
}

func NewLoginHandler(app *echo.Echo, db *db.Queries, ctx context.Context) *LoginHandler {
	return &LoginHandler{app, db, ctx}
}

func (h LoginHandler) HandleAllRoutes() {
	// pages
	h.app.GET("/login", h.HandleLoginShow)

	// apis
	h.app.POST("/api/login", h.HandleUserLogin)
}

func (h LoginHandler) HandleLoginShow(c echo.Context) error {
	return render(c, login.LoginShow())
}

func (h LoginHandler) HandleUserLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	data := login.NewLoginFormData()

	if email == "" {
		return render(c, login.LoginForm(login.ErrorNoEmail, *data))
	}
	if password == "" {
		data.Email = email
		return render(c, login.LoginForm(login.ErrorNoPassword, *data))
	}

	user, err := h.db.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		data.Email = email
		return render(c, login.LoginForm(login.ErrorEmailNotFound, *data))
	}

	ok := comparePasswordHash(password, user.Password)
	if !ok {
		data.Email = email
		return render(c, login.LoginForm(login.ErrorWrongPassword, *data))
	}

	// cookie := new(http.Cookie)

	return render(c, login.LoginForm(login.ErrorNone, *data))
}
