package adapters

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Binary Interface for interacting with the use case repository.
type Binary interface {
	CreateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error)
	UpdateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error)
	DeleteBinary(ctx context.Context, ID uuid.UUID) error

	BinaryReader
}

// BinaryReader user-readable interface
type BinaryReader interface {
	ListBinary(ctx context.Context, parameter queryParameter.QueryParameter) (*binary.ListBinaryViewModel, error)
	CountBinary(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
