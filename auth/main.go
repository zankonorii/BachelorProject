package main

import (
	"fmt"
	"jwt_client/database"
	"jwt_client/handler"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Starting Auth Service ", "time":"%s"}`, time.Now()))
	e := echo.New()

	db, err := database.NewMySQL()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	h := handler.NewHandler{
		DB: db,
	}
	
	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Registration routes ", "time":"%s"}`, time.Now()))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/check-login", h.CheckLogIn)
	e.POST("/login", h.LogIn)
	e.POST("/register", h.Register)

	e.Logger.Fatal(e.Start(":9001"))
}
