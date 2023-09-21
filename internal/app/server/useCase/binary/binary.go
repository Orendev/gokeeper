package binary

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(binaries ...*binary.BinaryData) ([]*binary.BinaryData, error) {
	return uc.adapterStorage.CreateBinary(binaries...)
}

func (uc *UseCase) Update(binaryUpdate binary.BinaryData) (*binary.BinaryData, error) {
	return uc.adapterStorage.UpdateBinary(binaryUpdate.ID(), func(oldBinary *binary.BinaryData) (*binary.BinaryData, error) {
		return binary.NewWithID(
			oldBinary.ID(),
			oldBinary.UserID(),
			binaryUpdate.Title(),
			binaryUpdate.Body(),
			binaryUpdate.Comment(),
			binaryUpdate.Version(),
			oldBinary.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteBinary(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
	return uc.adapterStorage.ListBinary(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *binary.BinaryData, err error) {
	return uc.adapterStorage.ReadBinaryByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountBinary()
}
