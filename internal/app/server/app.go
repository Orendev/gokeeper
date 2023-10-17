package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Orendev/gokeeper/internal/app/server/configs"
	deliveryGrpc "github.com/Orendev/gokeeper/internal/app/server/delivery/grpc"
)

// Run  starts the application.
func Run(delivery *deliveryGrpc.Delivery, cfg *configs.Config) {
	var err error
	var wg sync.WaitGroup
	ctx := gracefulShutdown()

	wg.Add(1)

	// create grpc server

	go func() {
		defer wg.Done()
		fmt.Printf("service started successfully on grpc port: %d", cfg.Delivery.GRPC.Port)
		// получаем запрос gRPC

		if err = delivery.Run(); err != nil {
			log.Fatalf("failed to start grpc server %s", err)
		}
	}()

	<-ctx.Done()

	delivery.ShutDown()
	wg.Wait()
}

func gracefulShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	irqSig := make(chan os.Signal, 1)
	// Получено сообщение о завершении работы от операционной системы.
	signal.Notify(irqSig, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		<-irqSig
		cancel()
	}()
	return ctx
}
