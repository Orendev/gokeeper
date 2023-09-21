package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Orendev/gokeeper/internal/app/client/configs"
	deliveryCLI "github.com/Orendev/gokeeper/internal/app/client/delivery/cli"
	repositorySQLite "github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite"
	useCaseUser "github.com/Orendev/gokeeper/internal/app/client/useCase/user"
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
	fmt.Println(cfg)
	repoStorageSQLite, err := repositorySQLite.New(
		conn.DB,
		cfg.SQLite,
	)
	if err != nil {
		log.Fatal(err)
	}

	var (
		ucUser      = useCaseUser.New(repoStorageSQLite, useCaseUser.Options{})
		deliveryCLI = deliveryCLI.New(ucUser)
	)

	//go func() {
	//	if err := deliveryCLI.Run(); err != nil {
	//		log.Fatalf("failed to start client %s", err)
	//	}
	//}()

	if err := deliveryCLI.Run(); err != nil {
		log.Fatalf("failed to start client %s", err)
	}

	//signalCh := make(chan os.Signal, 1)
	//signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	//<-signalCh
}
