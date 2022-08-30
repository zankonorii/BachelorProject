package main

import (
	"github.com/labstack/echo/v4/middleware"
	"zankonorii/products/service/handller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", handller.GetProducts)

	e.Logger.Fatal(e.Start(":8100"))
}
