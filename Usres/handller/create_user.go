package handller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserRequest struct{
	Name string `json:"name"`
	LastName string `json:"last_name"`
}

func (h NewHandller)CreateUser(ctx echo.Context) error{
	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Start CreateUser Handller ", "time":"%s"}`, time.Now()))
	 user := &UserRequest{}

	if err:= ctx.Bind(user); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Sended Data ", "time":"%s", "data":"{"name":"%s", "lastName":"%s"}"}`,
	 time.Now(), user.Name, user.LastName))

	sql := "INSERT INTO users(name, last_name, updated_at, created_at) VALUES(?,?,current_timestamp, current_timestamp)"
	stmt , err := h.DB.Prepare(sql)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.LastName)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Added User To DataBase ", "time":"%s", "data":"%s"}`, time.Now(), result))

	return ctx.JSON(http.StatusCreated, user)
}