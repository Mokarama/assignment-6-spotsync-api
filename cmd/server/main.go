package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Test Route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "SpotSync API is running successfully!",
		})
	})

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
