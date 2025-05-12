package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lMikadal/go-backend_api_user/internal/auth"
	"github.com/lMikadal/go-backend_api_user/internal/database"
	"github.com/lMikadal/go-backend_api_user/model"
	"github.com/labstack/echo/v4"
)

func (apiConfig *ApiConfig) HandlerCreateUser(c echo.Context) error {
	type User struct {
		Name string `json:"name"`
	}

	u := User{}
	if err := c.Bind(&u); err != nil {
		return c.String(http.StatusBadRequest, "something wrong!!")
	}

	user, err := apiConfig.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      u.Name,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "something wrong!!")
	}

	respUser := model.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}

	return c.JSON(http.StatusCreated, model.DatabaseUserToUser(respUser))
}

func (apiConfig *ApiConfig) HandlerGetUser(c echo.Context) error {
	apiKey, err := auth.GetAPIKey(c.Request().Header)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	user, err := apiConfig.DB.GetUserByAPIKey(c.Request().Context(), apiKey)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something wrong!!")
	}

	respUser := model.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}

	return c.JSON(http.StatusOK, model.DatabaseUserToUser(respUser))
}
