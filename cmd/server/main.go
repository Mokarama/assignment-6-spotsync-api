package main

import (
	"net/http"

	"github.com/Mokarama/assignment-6-spotsync-api/config"
	"github.com/Mokarama/assignment-6-spotsync-api/database"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to PostgreSQL database
	database.ConnectDatabase()

	// Create Echo instance
	e := echo.New()

	// Test Route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "SpotSync API is running successfully!",
		})
	})

	// Start Server
	e.Logger.Fatal(e.Start(":" + config.GetEnv("PORT")))
}
