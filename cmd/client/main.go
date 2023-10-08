package main

import (
	"log"

	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/grpc"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/grpc/interceptors"
	repositorySQLite "github.com/Orendev/gokeeper/internal/pkg/repository/storage/sqlite"
	useCaseAccountStorage "github.com/Orendev/gokeeper/internal/pkg/useCase/account"
	useCaseBinaryStorage "github.com/Orendev/gokeeper/internal/pkg/useCase/binary"
	useCaseCardStorage "github.com/Orendev/gokeeper/internal/pkg/useCase/card"
	useCaseTextStorage "github.com/Orendev/gokeeper/internal/pkg/useCase/text"
	useCaseUserStorage "github.com/Orendev/gokeeper/internal/pkg/useCase/user"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	memory "github.com/Orendev/gokeeper/pkg/tools/fileStorage"

	"github.com/Orendev/gokeeper/internal/app/client/configs"
	deliveryCLI "github.com/Orendev/gokeeper/internal/app/client/delivery/cli"
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

	conn, err := sqlite.New(cfg.SQLite)

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
		ucUserClient  = useCaseUserStorage.New(repoClient, useCaseUserStorage.Options{})

		ucAccountStorage = useCaseAccountStorage.New(repoSQLite, useCaseAccountStorage.Options{})
		ucAccountClient  = useCaseAccountStorage.New(repoClient, useCaseAccountStorage.Options{})

		ucTextStorage   = useCaseTextStorage.New(repoSQLite, useCaseTextStorage.Options{})
		ucTextClient    = useCaseTextStorage.New(repoClient, useCaseTextStorage.Options{})
		ucBinaryStorage = useCaseBinaryStorage.New(repoSQLite, useCaseBinaryStorage.Options{})
		ucBinaryClient  = useCaseBinaryStorage.New(repoClient, useCaseBinaryStorage.Options{})
		ucCardStorage   = useCaseCardStorage.New(repoSQLite, useCaseCardStorage.Options{})
		ucCardClient    = useCaseCardStorage.New(repoClient, useCaseCardStorage.Options{})
	)

	cli := deliveryCLI.New(
		ucUserStorage,
		ucUserClient,
		ucAccountStorage,
		ucAccountClient,
		ucTextStorage,
		ucTextClient,
		ucBinaryStorage,
		ucBinaryClient,
		ucCardStorage,
		ucCardClient,
		memory.NewFileStorage(cfg.File.FileStoragePath),
		cfg.Auth.CryptoKeyJWT,
	)

	if err := cli.Run(); err != nil {
		log.Fatalf("failed to start client %s", err)
	}
}
