package main

import (
	"fmt"
	"time"
	"zankonorii/products/service/database"
	"zankonorii/products/service/handller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Starting Post Service ", "time":"%s"}`, time.Now()))
	e := echo.New()

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Set Database ", "time":"%s"}`, time.Now()))
	db, err := database.NewMySQL()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Initial handller structure ", "time":"%s"}`, time.Now()))
	h := handller.NewHandller{
		DB: db,
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Register routes", "time":"%s"}`, time.Now()))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/posts", h.GetPosts)
	// e.GET("/post/:id", h.GetPosts)
	// e.GET("/post/:pid/user/:uid", h.GetPosts)
	e.POST("/post", h.CreatePost)
	// e.PUT("/post/:id", h.GetPosts)
	// e.DELETE("/post/:id", h.GetPosts)


	e.Logger.Fatal(e.Start(":8300"))
}
