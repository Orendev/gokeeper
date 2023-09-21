package useCase

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Binary interface {
	Create(binaries ...*binary.BinaryData) ([]*binary.BinaryData, error)
	Update(binary binary.BinaryData) (*binary.BinaryData, error)
	Delete(ID uuid.UUID) error

	BinaryReader
}

type BinaryReader interface {
	List(parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error)
	ReadByID(ID uuid.UUID) (response *binary.BinaryData, err error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
