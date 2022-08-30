package handller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h NewHandller) DeleteUser (ctx echo.Context) error{
	user := &UserRequest{}

	if err:= ctx.Bind(user); err != nil {
		return err
	}

	_, err := h.DB.Query(
		fmt.Sprintf(`DELETE FROM users WHERE id = %s;`, 
		ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, user)
}