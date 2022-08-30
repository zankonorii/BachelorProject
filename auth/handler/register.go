package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserRequest struct{
	Name string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}

type UserResponce struct{
	Name string `json:"name"`
	LastName string `json:"last_name"`
	JWT string `json:"jwt"`
}


func (h NewHandler)Register(ctx echo.Context) error {
	user := &UserRequest{}

	if err := ctx.Bind(user); err != nil {
		logrus.Error(err)
		return err
	}

	validToken, err :=  h.GetJWT(user.Password)
	if err != nil {
		logrus.Error(err)
		return err
	}

	postBody , _ := json.Marshal(map[string]string{
		"name" : user.Name,
		"last_name" : user.LastName,
		"jwt" : validToken,
	})

	responseBody := bytes.NewBuffer(postBody)

	res , err := http.Post("http://127.0.0.1:8200/create-user","application/json", responseBody)

	if err != nil {
		logrus.Error(err)
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
	   logrus.Error(err)
	   return err
	}

	ur := UserResponce{}
	if err := json.Unmarshal(body, &ur); err != nil{
		logrus.Error(err)
		return err
	}
	
	return ctx.JSON(http.StatusAccepted, ur)
}