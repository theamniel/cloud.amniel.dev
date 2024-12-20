package main

import (
	"context"
	"fmt"
	"net"

	"cloud.amniel.dev/services/multimedia/protocols"
	"cloud.amniel.dev/services/multimedia/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	app := fx.New(
		fx.Supply(sugar, koanf.New(".")),
		fx.Invoke(server),
	)
	app.Run()
}

// server sets up and starts the gRPC server.
func server(lc fx.Lifecycle, logger *zap.SugaredLogger, k *koanf.Koanf) {
	srv := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
				return err
			}

			protocols.RegisterMultimediaServiceServer(srv, &services.MultimediaService{})
			list, err := net.Listen("tcp", fmt.Sprintf(":%d", k.Int("multimedia.port")))
			if err != nil {
				return err
			}

			go func() {
				logger.Infof("Running Grpc on \"%s\"", list.Addr().String())
				logger.Info("Press CTRL-C to stop the application")
				if err := srv.Serve(list); err != nil {
					logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down multimedia-service...")
			srv.GracefulStop()
			return nil
		},
	})
}
