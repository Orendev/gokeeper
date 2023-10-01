package text

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, texts ...*text.TextData) ([]*text.TextData, error) {
	return uc.adapterStorage.CreateText(ctx, texts...)
}

func (uc *UseCase) Update(ctx context.Context, textUpdate text.TextData) (*text.TextData, error) {
	return uc.adapterStorage.UpdateText(ctx, textUpdate.ID(), func(oldText *text.TextData) (*text.TextData, error) {
		return text.NewWithID(
			oldText.ID(),
			oldText.UserID(),
			textUpdate.Title(),
			textUpdate.Data(),
			textUpdate.Comment(),
			oldText.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteText(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
	return uc.adapterStorage.ListText(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountText(ctx, parameter)
}
