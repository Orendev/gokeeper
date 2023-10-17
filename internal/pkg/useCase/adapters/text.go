package adapters

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Text interface {
	CreateText(ctx context.Context, text *text.TextData) (*text.TextData, error)
	UpdateText(ctx context.Context, text *text.TextData) (*text.TextData, error)
	DeleteText(ctx context.Context, ID uuid.UUID) error

	TextReader
}

type TextReader interface {
	ListText(ctx context.Context, parameter queryParameter.QueryParameter) (*text.ListTextViewModel, error)
	CountText(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
