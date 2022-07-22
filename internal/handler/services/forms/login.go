package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	formsEntity "github.com/fiber-go-sis-app/internal/entity/forms"

	customPkg "github.com/fiber-go-sis-app/utils/pkg/custom"

	formsSvc "github.com/fiber-go-sis-app/internal/usecase/services/forms"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var loginRequest formsEntity.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validator.New().Struct(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := formsSvc.LoginForm(ctx, loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return customPkg.BuildJSONRes(ctx, data)
}
