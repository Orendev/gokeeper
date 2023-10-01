package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Text interface {
	CreateText(ctx context.Context, texts ...*text.TextData) ([]*text.TextData, error)
	UpdateText(ctx context.Context, ID uuid.UUID, updateFn func(t *text.TextData) (*text.TextData, error)) (*text.TextData, error)
	DeleteText(ctx context.Context, ID uuid.UUID) error

	TextReader
}

type TextReader interface {
	ListText(ctx context.Context, parameter queryParameter.QueryParameter) ([]*text.TextData, error)
	CountText(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
