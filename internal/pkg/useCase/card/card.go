package card

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, card *card.CardData) (*card.CardData, error) {
	return uc.adapterStorage.CreateCard(ctx, card)
}

func (uc *UseCase) Update(ctx context.Context, card *card.CardData) (*card.CardData, error) {
	return uc.adapterStorage.UpdateCard(ctx, card)
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteCard(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) (*card.ListCardViewModel, error) {
	return uc.adapterStorage.ListCard(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountCard(ctx, parameter)
}
