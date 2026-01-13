package controller

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/swagger/v2"
	"github.com/spf13/viper"
	"github.com/yockii/inkruins/internal/controller/docs"
)

func init() {
	viper.SetDefault("server.swaggerEnable", false)
}

type Controller interface {
	InitializeRouter(router fiber.Router)
}

var controllers []Controller

// @title InkRuins API
// @version 1.0
// @description 墨墟API文档
// @host localhost:8080
// @BasePath ./
// @securityDefinitions.basic  BasicAuth

func InitializeRouter(app *fiber.App) {
	app.Get("/health", healthcheck.New())

	if viper.GetBool("server.swaggerEnable") {
		slog.Info("swagger enabled", "swagger", docs.SwaggerInfo)
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	app.Use(recoverer.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())

	router := app.Group("/api/v1")
	// 初始化所有需要认证的controller
	for _, c := range controllers {
		c.InitializeRouter(router)
	}
}
