package handller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PostResponse struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Likes     int    `json:"likes"`
	Views     int    `json:"views"`
	Image     string `json:"image"`
	Available bool   `json:"available"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func (h NewHandller) GetPosts(ctx echo.Context) error {
	posts := []PostResponse{}

	results, err := h.DB.Query("SELECT * FROM posts")
	if err != nil {
		logrus.Error(err)
	}

	for results.Next(){
		var post PostResponse

		err = results.Scan(&post.ID, &post.UserId, &post.Likes, &post.Views, &post.Image, &post.Available, &post.Name, &post.UpdatedAt, &post.CreatedAt)

		if err != nil {
			logrus.Error(err)
		}

		posts = append(posts, post)
	}

	return ctx.JSON(http.StatusOK, posts)

}

func (h NewHandller) GetPost(ctx echo.Context) error {
	var post PostResponse

	err := h.DB.QueryRow("SELECT * FROM posts WHERE id = ?", ctx.Param("id")).Scan(&post.ID, &post.UserId, &post.Likes, &post.Views, &post.Image, &post.Available, &post.Name, &post.UpdatedAt, &post.CreatedAt)
	if err != nil {
		logrus.Error(err)
	}
	return ctx.JSON(http.StatusAccepted, post)
}
