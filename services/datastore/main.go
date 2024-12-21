package main

import (
	"context"
	"fmt"
	"net"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"cloud.amniel.dev/services/datastore/services/database"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	k := koanf.New(".")
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		sugar.Fatalf("error loading config: %v", err)
	}

	fx.New(
		fx.Supply(sugar, k),
		fx.Provide(database.New),
		fx.Invoke(server),
	).Run()
}

func server(lc fx.Lifecycle, k *koanf.Koanf, logger *zap.SugaredLogger) {
	srv := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			list, err := net.Listen("tcp", fmt.Sprintf(":%d", k.Int("datastore.port")))
			if err != nil {
				return err
			}
			go func() {
				logger.Infof("Running Grpc on \"%s\"", list.Addr().String())
				logger.Info("Press CTRL-C to stop the application.")
				if err := srv.Serve(list); err != nil {
					logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down datastore-service...")
			srv.GracefulStop()
			return nil
		},
	})
}
