package main

import (
	"net/http"

	"github.com/Mokarama/assignment-6-spotsync-api/config"
	"github.com/Mokarama/assignment-6-spotsync-api/database"
	"github.com/Mokarama/assignment-6-spotsync-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect database
	database.ConnectDatabase()

	// Create Echo instance
	e := echo.New()

	// Health Check Route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "SpotSync API is running successfully!",
		})
	})

	// Register Routes
	routes.RegisterAuthRoutes(e)

	// Start Server
	e.Logger.Fatal(e.Start(":" + config.GetEnv("PORT")))
}
