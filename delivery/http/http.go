package http

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"github.com/rizama/favorite-book-tracker/delivery/http/router"
	"github.com/rizama/favorite-book-tracker/domain"
)

func NewHttpDelivery(domain domain.Domain, engine *html.Engine) *fiber.App {
	config := fiber.Config{
		AppName:           os.Getenv("APP_NAME"),
		EnablePrintRoutes: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		Views:             engine,
	}

	if os.Getenv("GO_ENV") == "production" {
		config.Prefork = true
	}

	app := fiber.New(config)
	app.Use(logger.New())
	app.Static("/static", "resource")
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	// web route
	router.NewRouter(app, domain)

	return app
}
