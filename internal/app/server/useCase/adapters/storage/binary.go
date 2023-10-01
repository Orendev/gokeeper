package storage

import (
	"context"

	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Binary interface {
	CreateBinary(ctx context.Context, binaries ...*binary.BinaryData) ([]*binary.BinaryData, error)
	UpdateBinary(ctx context.Context, ID uuid.UUID, updateFn func(b *binary.BinaryData) (*binary.BinaryData, error)) (*binary.BinaryData, error)
	DeleteBinary(ctx context.Context, ID uuid.UUID) error

	BinaryReader
}

type BinaryReader interface {
	ListBinary(ctx context.Context, parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error)
	CountBinary(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error)
}
