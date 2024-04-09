package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHealthCheck(c echo.Context) error {
	message := fmt.Sprintf("Hello! you've requested: %s", c.Path())
	return c.JSON(
		http.StatusOK,
		healthCheckResponse{
			Message: message,
		},
	)
}

type healthCheckResponse struct {
	Message string `json:"message"`
}
