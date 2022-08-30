package handller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PostRequest struct {
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Available bool   `json:"available"`
}

func (h NewHandller) CreatePost(ctx echo.Context) error {
	post := &PostRequest{}
	if err := ctx.Bind(post); err != nil {
		logrus.Error(err)
		return err
	}

	sql := "INSERT INTO posts(user_id, name, image, available, likes, views, updated_at, created_at) VALUES(?,?,?,?,0,0,current_timestamp, current_timestamp)"
	stmt, err := h.DB.Prepare(sql)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(post.UserId, post.Name, post.Image, post.Available)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Added User To DataBase ", "time":"%s", "data":"%s"}`, time.Now(), result))

	return ctx.JSON(http.StatusCreated, post)
}
