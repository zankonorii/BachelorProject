package handller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h NewHandller) Updatepost (ctx echo.Context) error {
	post := &PostRequest{}

	if err:= ctx.Bind(post); err != nil {
		return err
	}

	_, err := h.DB.Query(
		fmt.Sprintf(`UPDATE posts SET name="%s", image ="%s", available=%t, updated_at = CURRENT_TIMESTAMP WHERE id = %s;`, 
		post.Name, post.Image, post.Available, ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, post)
}