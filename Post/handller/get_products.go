package handller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Image     string  `json:"image"`
	Available bool    `json:"available"`
}

func GetProducts(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, getProducts())
}

func getProducts() []ProductResponse {
	return []ProductResponse{
		{1, "One", 2.5, "https://i.pravatar.cc/300", true},
		{2, "Three", 0.5, "https://i.pravatar.cc/300", true},
		{3, "Four", 1.5, "https://i.pravatar.cc/300", false},
		{4, "Five", 5.5, "https://i.pravatar.cc/300", true},
		{5, "Six", 6.0, "https://i.pravatar.cc/300", false},
		{6, "Seven", 3.2, "https://i.pravatar.cc/300", true},
	}
}
