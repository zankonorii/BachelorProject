package handller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h NewHandller)IncreaseLike(ctx echo.Context) error {
	_, err := h.DB.Query(
		fmt.Sprintf(`UPDATE posts SET likes = likes + 1 , updated_at = CURRENT_TIMESTAMP WHERE id = %s;`, 
		ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, "Increased successfully")
}

func (h NewHandller)DecreaseLike(ctx echo.Context) error {
	_, err := h.DB.Query(
		fmt.Sprintf(`UPDATE posts SET likes = likes - 1 , updated_at = CURRENT_TIMESTAMP WHERE id = %s;`, 
		ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, "Decreased successfully")
}

func (h NewHandller)IncreaseView(ctx echo.Context) error {
	_, err := h.DB.Query(
		fmt.Sprintf(`UPDATE posts SET views = views + 1 , updated_at = CURRENT_TIMESTAMP WHERE id = %s;`, 
		ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, "Increased successfully")
}

func (h NewHandller)DecreaseView(ctx echo.Context) error {
	_, err := h.DB.Query(
		fmt.Sprintf(`UPDATE posts SET views = views - 1 , updated_at = CURRENT_TIMESTAMP WHERE id = %s;`, 
		ctx.Param("id")))
	if err != nil {
		logrus.Error(err)
		return err
	}	
	return ctx.JSON(http.StatusAccepted, "Decreased successfully")
}
