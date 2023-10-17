package useCase

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Binary Interface for interaction between delivery and use case
type Binary interface {
	Create(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error)
	Update(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error)
	Delete(ctx context.Context, ID uuid.UUID) error

	BinaryReader
}

// BinaryReader user-readable interface
type BinaryReader interface {
	List(ctx context.Context, parameter queryParameter.QueryParameter) (*binary.ListBinaryViewModel, error)
	Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
