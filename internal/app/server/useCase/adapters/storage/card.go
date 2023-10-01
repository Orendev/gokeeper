package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Card interface {
	CreateCard(ctx context.Context, cards ...*card.CardData) ([]*card.CardData, error)
	UpdateCard(ctx context.Context, ID uuid.UUID, updateFn func(c *card.CardData) (*card.CardData, error)) (*card.CardData, error)
	DeleteCard(ctx context.Context, ID uuid.UUID) error

	CardReader
}

type CardReader interface {
	ListCard(ctx context.Context, parameter queryParameter.QueryParameter) ([]*card.CardData, error)
	CountCard(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
