package storage

import "github.com/Orendev/gokeeper/internal/pkg/useCase/adapters"

// Interface for interacting with the use case repository.

type Storage struct {
	User
	adapters.Text
	adapters.Binary
	adapters.Card
	Account
}
