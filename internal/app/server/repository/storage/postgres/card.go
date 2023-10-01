package postgres

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

var tableNameCard = "cards"

func (r *Repository) CreateCard(ctx context.Context, cards ...*card.CardData) ([]*card.CardData, error) {
	panic("implement me")
}

func (r *Repository) UpdateCard(ctx context.Context, ID uuid.UUID, updateFn func(a *card.CardData) (*card.CardData, error)) (*card.CardData, error) {
	panic("implement me")
}

func (r *Repository) DeleteCard(ctx context.Context, ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListCard(ctx context.Context, parameter queryParameter.QueryParameter) ([]*card.CardData, error) {
	panic("implement me")
}

func (r *Repository) CountCard(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	panic("implement me")
}
