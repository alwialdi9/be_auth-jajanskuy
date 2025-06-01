package middlewares

import (
	"strings"

	"github.com/alwialdi9/be_auth-jajanskuy/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check if the request has an Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.JsonErrorResponse(c, fiber.StatusUnauthorized, "Authorization header is missing", nil)
	}

	// Validate the token (this is a placeholder, implement your token validation logic)
	checkToken := authHeader[len("Bearer "):]
	if checkToken == "" {
		return utils.JsonErrorResponse(c, fiber.StatusUnauthorized, "Invalid token", nil)
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenStr, &utils.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return utils.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims := token.Claims.(*utils.Claims)
	c.Locals("user", claims)
	c.Locals("user_id", claims.UserID)
	c.Locals("username", claims.Username)
	c.Locals("email", claims.Email)
	return c.Next()
}
