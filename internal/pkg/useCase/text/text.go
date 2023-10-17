package text

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, text *text.TextData) (*text.TextData, error) {
	return uc.adapterStorage.CreateText(ctx, text)
}

func (uc *UseCase) Update(ctx context.Context, text *text.TextData) (*text.TextData, error) {
	return uc.adapterStorage.UpdateText(ctx, text)
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteText(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) (*text.ListTextViewModel, error) {
	return uc.adapterStorage.ListText(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountText(ctx, parameter)
}
