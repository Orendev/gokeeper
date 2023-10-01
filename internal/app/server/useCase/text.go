package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Text interface {
	Create(ctx context.Context, texts ...*text.TextData) ([]*text.TextData, error)
	Update(ctx context.Context, text text.TextData) (*text.TextData, error)
	Delete(ctx context.Context, ID uuid.UUID) error

	TextReader
}

type TextReader interface {
	List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*text.TextData, error)
	Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
