package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Request struct {
	Name string `json:"name"`
	JWT string `json:"jwt"`
}

func (h NewHandler)CheckLogIn(ctx echo.Context) error {
	r := &Request{}
	user := UserResponse{}

	if err := ctx.Bind(r); err != nil {
		logrus.Error(err)
		return err
	}
	if err := h.DB.QueryRow("SELECT * FROM users WHERE name = ?", r.Name).Scan(&user.ID, &user.Name, &user.LastName, &user.JWT, &user.UpdatedAt, &user.CreatedAt); err != nil {
		logrus.Error(err)
		return ctx.JSON(http.StatusUnauthorized, "please login or register")

	}

	if user.JWT != r.JWT {
		return ctx.JSON(http.StatusUnauthorized, "please login or register")
	}

	return ctx.JSON(http.StatusAccepted, user)
}