package handller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h NewHandller) EditUser (ctx echo.Context) error{
	user := &UserRequest{}

	if err:= ctx.Bind(user); err != nil {
		return err
	}

	_, err := h.DB.Query(
		fmt.Sprintf(`UPDATE users SET name="%s", last_name ="%s", updated_at = CURRENT_TIMESTAMP WHERE id = %s;`, 
		user.Name, user.LastName, ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, user)
}