package handlers

import (
	"context"
	"log"
	"net/mail"

	"github.com/jacksonopp/tanglefit/db"
	"github.com/jacksonopp/tanglefit/view/components"
	"github.com/jacksonopp/tanglefit/view/signup"
	"github.com/labstack/echo/v4"
)

type SignupHandler struct {
	app *echo.Echo
	db  *db.Queries
	ctx context.Context
}

func NewSignupHandler(app *echo.Echo, db *db.Queries, ctx context.Context) *SignupHandler {
	return &SignupHandler{app, db, ctx}
}

func (h SignupHandler) HandleAllRoutes() {
	// pages
	h.app.GET("/signup", h.HandleSignupShow)

	// api
	h.app.POST("/api/signup", h.HandleSignup)
	h.app.POST("/api/signup/validate-email", h.HandleValidateEmail)
}

func (h SignupHandler) HandleSignupShow(c echo.Context) error {
	return render(c, signup.SignUpShow())
}

func (h SignupHandler) HandleSignup(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")

	data := signup.NewSignUpFormData(signup.WithEmail(email))

	log.Println(email, password, confirmPassword)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return render(c, signup.SignUpForm(signup.ErrorInvalidEmail, *data))
	}

	if password != confirmPassword {
		return render(c, signup.SignUpForm(signup.ErrorPasswordMismatch, *data))
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Println("ERROR:", err)
		return render(c, signup.SignUpForm(signup.ErrorUnknown, *data))
	}

	err = h.db.CreateUser(h.ctx, db.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
		Role: db.NullRole{
			Role:  db.RoleMember,
			Valid: true,
		},
	})
	if err != nil {
		log.Println("ERROR:", err)
		return render(c, signup.SignUpForm(signup.ErrorUnknown, *data))
	}

	return render(c, signup.SignUpForm(signup.Success, *data))
}

func (h SignupHandler) HandleValidateEmail(c echo.Context) error {
	email := c.FormValue("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		return render(c, components.ErrorMessage("Email must be valid (ex: me@example.com)"))
	}
	return render(c, components.ErrorMessage(""))
}
