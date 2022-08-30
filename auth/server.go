package main

import (
	"zankonorii/auth/service/handller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/login", handller.Login)
	//e.GET("/register", handller.GetProducts)
	//e.GET("/logout", handller.GetProducts)

	e.Logger.Fatal(e.Start(":1525"))
}
