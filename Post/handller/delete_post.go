package handller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h NewHandller)DeletePost(ctx echo.Context) error {
	post := &PostRequest{}

	if err:= ctx.Bind(post); err != nil {
		return err
	}

	_, err := h.DB.Query(
		fmt.Sprintf(`DELETE FROM posts WHERE id = %s;`, 
		ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, post)
}