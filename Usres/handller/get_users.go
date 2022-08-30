package handller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UpdatedAt    string    `json:"updated_at"`
	CreatedAt    string    `json:"created_at"`
}

func (h NewHandller)GetUsers(ctx echo.Context) error {
	
	users := []UserResponse{}

	results, err := h.DB.Query("SELECT * FROM users")
	if err != nil {
		logrus.Error(err)
	}

	for results.Next(){
		var user UserResponse

		err = results.Scan(&user.ID, &user.Name, &user.LastName , &user.UpdatedAt, &user.CreatedAt)

		if err != nil {
			logrus.Error(err)
		}

		users = append(users, user)
	}

	return ctx.JSON(http.StatusOK, users)
}

func (h NewHandller) GetUser (ctx echo.Context) error {

	var user UserResponse

	err := h.DB.QueryRow("SELECT * FROM users WHERE id = ?", ctx.Param("id")).Scan(&user.ID, &user.Name, &user.LastName , &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		logrus.Error(err)
	}
	return ctx.JSON(http.StatusAccepted, user)
}
