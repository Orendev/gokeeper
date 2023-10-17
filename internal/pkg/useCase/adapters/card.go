package adapters

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Card Interface for interacting with the use case repository.
type Card interface {
	CreateCard(ctx context.Context, card *card.CardData) (*card.CardData, error)
	UpdateCard(ctx context.Context, card *card.CardData) (*card.CardData, error)
	DeleteCard(ctx context.Context, ID uuid.UUID) error

	CardReader
}

// CardReader user-readable interface
type CardReader interface {
	ListCard(ctx context.Context, parameter queryParameter.QueryParameter) (*card.ListCardViewModel, error)
	CountCard(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
