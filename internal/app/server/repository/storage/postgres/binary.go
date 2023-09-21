package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

func (r *Repository) CreateBinary(binaries ...*binary.BinaryData) ([]*binary.BinaryData, error) {
	panic("implement me")
}

func (r *Repository) UpdateBinary(ID uuid.UUID, updateFn func(a *binary.BinaryData) (*binary.BinaryData, error)) (*binary.BinaryData, error) {
	panic("implement me")
}

func (r *Repository) DeleteBinary(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListBinary(parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
	panic("implement me")
}

func (r *Repository) ReadBinaryByID(ID uuid.UUID) (response *binary.BinaryData, err error) {
	panic("implement me")
}

func (r *Repository) CountBinary() (uint64, error) {
	panic("implement me")
}
