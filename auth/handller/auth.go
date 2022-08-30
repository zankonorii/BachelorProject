package handller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(ctx echo.Context)  error{
	return ctx.JSON(http.StatusOK, "you are login")
}
