package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Card Interface for interaction between delivery and use case
type Card interface {
	Create(ctx context.Context, card *card.CardData) (*card.CardData, error)
	Update(ctx context.Context, card *card.CardData) (*card.CardData, error)
	Delete(ctx context.Context, ID uuid.UUID) error

	CardReader
}

// CardReader user-readable interface
type CardReader interface {
	List(ctx context.Context, parameter queryParameter.QueryParameter) (*card.ListCardViewModel, error)
	Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
