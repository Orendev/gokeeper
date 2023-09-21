package storage

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Binary interface {
	CreateBinary(binaries ...*binary.BinaryData) ([]*binary.BinaryData, error)
	UpdateBinary(ID uuid.UUID, updateFn func(b *binary.BinaryData) (*binary.BinaryData, error)) (*binary.BinaryData, error)
	DeleteBinary(ID uuid.UUID) error

	BinaryReader
}

type BinaryReader interface {
	ListBinary(parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error)
	ReadBinaryByID(ID uuid.UUID) (response *binary.BinaryData, err error)
	CountBinary( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
