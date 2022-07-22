package personalia

import (
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
)

func WebSISUserHandler(ctx *fiber.Ctx) error {

	return ctx.Render("templates/sis/pages/user", constantsEntity.WebData{
		Title:        constantsEntity.WebSISUserTitle,
		BaseURL:      constantsEntity.BaseURL,
		StaticUrl:    constantsEntity.StaticUrl,
		CurrentURL:   constantsEntity.WebSISUserURL,
		LinkPageList: constantsEntity.LinkPageList,
	})
}
