package main

import (
	"fmt"
	"time"

	"zankonorii/users/service/database"
	"zankonorii/users/service/handller"

	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func main() {
	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Starting Project ", "time":"%s"}`, time.Now()))
	e  := echo.New()

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Set Up DataBase ", "time":"%s"}`, time.Now()))
	db, err := database.NewMySQL() 
	if err!= nil {
		panic(err)
	}

	h := handller.NewHandller{
		DB: db,
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Register Routes ", "time":"%s"}`, time.Now()))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/users", h.GetUsers)
	e.GET("/user/:id", h.GetUser)
	e.PUT("/user/:id", h.EditUser)
	e.POST("/create-user", h.CreateUser)
	e.DELETE("/user/:id", h.DeleteUser)

	e.Logger.Fatal(e.Start(":8200"))
}