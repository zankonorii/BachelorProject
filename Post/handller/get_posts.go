package handller

import (
	"github.com/labstack/echo/v4"
	"net/http"
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

	return ctx.JSON(http.StatusOK, nil)
}
