package handler

import (
	"net/http"

	"github.com/lMikadal/go-backend_api_user/internal/database"
	"github.com/labstack/echo/v4"
)

type ApiConfig struct {
	DB *database.Queries
}

func HandlerHealthy(c echo.Context) error {
	return c.String(http.StatusOK, "ok!!")
}
