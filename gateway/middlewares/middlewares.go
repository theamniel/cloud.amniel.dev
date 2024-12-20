package middlewares

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func Configure(app *fiber.App, zapLogger *zap.SugaredLogger) {
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: zapLogger.Desugar(),
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(recover.New())
}
