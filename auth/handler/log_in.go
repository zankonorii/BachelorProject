package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LogInInfo struct{
	Name string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UpdatedAt    string    `json:"updated_at"`
	CreatedAt    string    `json:"created_at"`
}

func (h NewHandler)LogIn(ctx echo.Context) error {
	info := &LogInInfo{}
	user := UserResponse{}
	var UJWT string

	if err := ctx.Bind(info); err != nil {
		logrus.Error(err)
		return err
	}

	validToken, err :=  h.GetJWT(info.Password)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if err := h.DB.QueryRow("SELECT * FROM users WHERE name = ?", info.Name).Scan(&user.ID, &user.Name, &user.LastName, &UJWT, &user.UpdatedAt, &user.CreatedAt); err != nil {
		logrus.Error(err)
		return ctx.JSON(http.StatusUnauthorized , "please login or register")

	}
	

	if UJWT != validToken {
		return ctx.JSON(http.StatusUnauthorized , "please login or register")
	}

	return ctx.JSON(http.StatusUnauthorized , "please login or register")

}