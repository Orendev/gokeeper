package card

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, cards ...*card.CardData) ([]*card.CardData, error) {
	return uc.adapterStorage.CreateCard(ctx, cards...)
}

func (uc *UseCase) Update(ctx context.Context, cardUpdate card.CardData) (*card.CardData, error) {
	return uc.adapterStorage.UpdateCard(ctx, cardUpdate.ID(), func(oldCard *card.CardData) (*card.CardData, error) {
		return card.NewWithID(
			oldCard.ID(),
			oldCard.UserID(),
			cardUpdate.CardNumber(),
			cardUpdate.CardName(),
			cardUpdate.CVC(),
			cardUpdate.CardDate(),
			cardUpdate.Comment(),
			oldCard.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteCard(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*card.CardData, error) {
	return uc.adapterStorage.ListCard(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountCard(ctx, parameter)
}
