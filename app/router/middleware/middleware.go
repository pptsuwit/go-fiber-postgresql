package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	jwtware "github.com/gofiber/jwt/v2"
)

func Authorize() fiber.Handler {

	return jwtware.New(jwtware.Config{
		// Filter:         nil,
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(viper.GetString("app.jwtSecret")),
		// SigningKeys:   nil,
		SigningMethod: "HS256",
		// ContextKey:    nil,
		// Claims:        nil,
		// TokenLookup:   nil,
		// AuthScheme:    nil,
	})
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}
func AuthError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   "Unauthorized",
		"message": err.Error(),
	})
	return nil
}
