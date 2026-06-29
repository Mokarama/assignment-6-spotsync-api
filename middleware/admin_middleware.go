package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		role := claims["role"].(string)

		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"success": false,
				"message": "Access denied. Admin only.",
			})
		}

		return next(c)
	}
}
