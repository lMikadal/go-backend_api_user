package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lMikadal/go-backend_api_user/handler"
	"github.com/lMikadal/go-backend_api_user/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq" // postgres driver
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// start use echo
	e := echo.New()

	// config CORS local
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// connect to database
	_, err = sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// group route
	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/healthy", handler.HandlerHealthy)

	e.Logger.Fatal(e.Start(":" + port))
}
