#!/usr/bin/make
SHELL = /bin/sh

APP_NAME = gokeeper

protoc:
	./scripts/gen-proto.sh


gen-cert:
	./scripts/gen-cert.sh


build-client: ## Build App Client
	go mod tidy
	go build -v -o $(APP_NAME)  -ldflags="-X 'github.com/Orendev/gokeeper/internal/app/client/delivery/cli.version=0.0.2'" ./internal/app/client/cmd/client/main.go
