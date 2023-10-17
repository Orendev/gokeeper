#!/usr/bin/make
SHELL = /bin/sh

APP_NAME = ./cmd/client/gokeeper
BUILD_VERSION = 0.0.2
BUILD_DATE = 2023-10-08
OS = linux windows darwin
ARCH_ALL=amd64 arm64

protoc:
	./scripts/gen-proto.sh


gen-cert:
	./scripts/gen-cert.sh


build-client: ## Build App Client
	go mod tidy
	./scripts/build.sh

mockery:
	./scripts/mockery.sh