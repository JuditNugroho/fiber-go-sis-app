package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	goccyJson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"

	"github.com/fiber-go-sis-app/utils/pkg/custom"
	"github.com/fiber-go-sis-app/utils/pkg/databases/elasticsearch"
	"github.com/fiber-go-sis-app/utils/pkg/databases/postgres"
	"github.com/fiber-go-sis-app/utils/pkg/jwt"

	serviceRoutes "github.com/fiber-go-sis-app/routes/services"
	webRoutes "github.com/fiber-go-sis-app/routes/web"
)

// Embed a template directory
//go:embed templates/*
var embedDirTemplate embed.FS

// Embed a static directory
//go:embed static/*
var embedDirStatic embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(embedDirTemplate), ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:        engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		JSONEncoder:  goccyJson.Marshal,
		JSONDecoder:  goccyJson.Unmarshal,
		AppName:      constantsEntity.AppName,
	})

	// Setting JWT RS256
	if err := jwt.GenerateJWT(); err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// Setting basic configuration
	app.Use(logger.New(), recover.New())
	app.Use(constantsEntity.StaticUrl, filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse:     true,
	}))

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := postgres.OpenConnection(); err != nil {
		panic(err)
	}

	if err := elasticsearch.NewESClient(); err != nil {
		panic(err)
	}

	if err := custom.SetupSchema(); err != nil {
		panic(err)
	}

	// Web handler login
	webRoutes.BuildLoginRoutes(app)

	// Web handler - SIS
	sisGroup := app.Group("/sis")
	webRoutes.BuildSISRoutes(sisGroup)

	// Service Group
	svcGroup := app.Group("/svc")
	serviceRoutes.BuildUserRoutes(svcGroup)
	serviceRoutes.BuildLoginRoutes(svcGroup)
	serviceRoutes.BuildMemberRoutes(svcGroup)
	serviceRoutes.BuildProductRoutes(svcGroup)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
