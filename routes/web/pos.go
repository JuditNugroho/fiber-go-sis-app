package web

import (
	"github.com/gofiber/fiber/v2"

	homeWeb "github.com/fiber-go-sis-app/internal/handler/web/pos/home"
)

func BuildPOSRoutes(service fiber.Router) {
	service.Get("/", homeWeb.WebPOSHomeHandler)
}
