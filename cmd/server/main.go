package main

import (
	"context"
	"log"
	"time"

	"github.com/Orendev/gokeeper/internal/app/server"
	"github.com/Orendev/gokeeper/internal/app/server/configs"
	deliveryGrpc "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc"
	repositoryStorage "github.com/Orendev/gokeeper/internal/pkg/repository/storage/postgres"
	useCaseAccount "github.com/Orendev/gokeeper/internal/pkg/useCase/account"
	useCaseBinary "github.com/Orendev/gokeeper/internal/pkg/useCase/binary"
	useCaseCard "github.com/Orendev/gokeeper/internal/pkg/useCase/card"
	useCaseText "github.com/Orendev/gokeeper/internal/pkg/useCase/text"
	useCaseUser "github.com/Orendev/gokeeper/internal/pkg/useCase/user"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/store/postgres"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
)

func main() {

	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := logger.NewLogger(cfg.Log.LogLevel); err != nil {
		log.Fatal(err)
	}

	conn, err := postgres.New(context.Background(), cfg.Postgres)
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	repoStorage, err := repositoryStorage.New(conn.Pool, cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	jwtManager := auth.NewJWTManager(cfg.Auth.CryptoKeyJWT, time.Duration(cfg.Auth.TokenDurationJWT)*time.Second)

	var (
		ucAccount    = useCaseAccount.New(repoStorage, useCaseAccount.Options{})
		ucUser       = useCaseUser.New(repoStorage, useCaseUser.Options{})
		ucCard       = useCaseCard.New(repoStorage, useCaseCard.Options{})
		ucText       = useCaseText.New(repoStorage, useCaseText.Options{})
		ucBinary     = useCaseBinary.New(repoStorage, useCaseBinary.Options{})
		deliveryGRPC = deliveryGrpc.New(ucUser, ucAccount, ucCard, ucText, ucBinary, jwtManager, cfg.Delivery.GRPC)
	)

	server.Run(deliveryGRPC, cfg)

}
