package middleware

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == constantsEntity.ErrMissingOrMalformedJWT {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrMissingOrMalformedJWT, "data": nil})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": constantsEntity.ErrInvalidORExpiredJWT, "data": nil})
}
