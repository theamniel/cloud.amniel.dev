package main

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"cloud.amniel.dev/gateway/middlewares"
	"cloud.amniel.dev/gateway/routes"
	"cloud.amniel.dev/utils"

	_ "github.com/asaskevich/govalidator"
	_ "google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	fx.New(
		fx.Supply(sugar, koanf.New(".")),
		fx.Provide(configureApp),
		fx.Invoke(
			middlewares.Configure,
			routes.Configure,
			server,
		),
	).Run()
}

func server(lc fx.Lifecycle, k *koanf.Koanf, logger *zap.SugaredLogger, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
				return err
			}

			id := utils.RandomString(32)
			hash := utils.Base64([]byte(id))
			if err := k.Set("gateway.token", hash); err != nil {
				return err
			}

			app.Hooks().OnListen(func(ld fiber.ListenData) error {
				if !fiber.IsChild() {
					scheme := "http"
					if ld.TLS {
						scheme += "s"
					}
					logger.Infof("Running server on \"%s://%s:%s\"", scheme, ld.Host, ld.Port)
					logger.Info("Press CTRL-C to stop the application")
				}
				return nil
			})

			go func() { // Never run hook `OnListen` if don't do this.
				if err := app.Listen(fmt.Sprintf(":%d", k.Int("gateway.port"))); err != nil {
					logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if !fiber.IsChild() {
				logger.Info("Shutting down connection...")
			}
			return app.ShutdownWithContext(ctx)
		},
	})
}

func configureApp(k *koanf.Koanf) *fiber.App {
	appConfig := fiber.Config{
		DisableStartupMessage: true,
		Prefork:               k.Bool("gateway.prefork"),
		BodyLimit:             k.Int("gateway.body_limit") << 20,
		ErrorHandler:          routes.ErrorHandler,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	}
	return fiber.New(appConfig)
}
