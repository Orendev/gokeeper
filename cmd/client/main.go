package main

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/interceptors"
	useCaseUserClient "github.com/Orendev/gokeeper/internal/app/client/useCase/client/user"
	useCaseUserStorage "github.com/Orendev/gokeeper/internal/app/client/useCase/storage/user"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"log"

	"github.com/Orendev/gokeeper/internal/app/client/configs"
	deliveryCLI "github.com/Orendev/gokeeper/internal/app/client/delivery/cli"
	repositorySQLite "github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/store/sqlite"
)

func main() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := logger.NewLogger(cfg.Log.LogLevel); err != nil {
		log.Fatal(err)
	}

	conn, err := sqlite.New(context.Background(), cfg.SQLite)

	if err != nil {
		panic(err)
	}

	defer func() {
		err := conn.DB.Close()
		if err != nil {

		}
	}()

	repoSQLite, err := repositorySQLite.New(
		conn.DB,
		cfg.SQLite,
	)
	if err != nil {
		log.Fatal(err)
	}

	authInterceptor, err := interceptors.NewAuthInterceptor(
		auth.AccessibleRoles(),
		&cfg.ServerGRPC)

	if err != nil {
		log.Fatal(err)
	}

	repoClient, err := grpc.New(authInterceptor, cfg.ServerGRPC)

	if err != nil {
		log.Fatal(err)
	}

	var (
		ucUserStorage = useCaseUserStorage.New(repoSQLite, useCaseUserStorage.Options{})
		ucUserClient  = useCaseUserClient.New(repoClient, useCaseUserClient.Options{})
		cli           = deliveryCLI.New(ucUserStorage, ucUserClient)
	)

	if err := cli.Run(); err != nil {
		log.Fatalf("failed to start client %s", err)
	}
}
