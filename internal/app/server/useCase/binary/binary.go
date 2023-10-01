package binary

import (
	"context"
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, binaries ...*binary.BinaryData) ([]*binary.BinaryData, error) {
	return uc.adapterStorage.CreateBinary(ctx, binaries...)
}

func (uc *UseCase) Update(ctx context.Context, binaryUpdate binary.BinaryData) (*binary.BinaryData, error) {
	return uc.adapterStorage.UpdateBinary(ctx, binaryUpdate.ID(), func(oldBinary *binary.BinaryData) (*binary.BinaryData, error) {
		return binary.NewWithID(
			oldBinary.ID(),
			oldBinary.UserID(),
			binaryUpdate.Title(),
			binaryUpdate.Data(),
			binaryUpdate.Comment(),
			oldBinary.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteBinary(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
	return uc.adapterStorage.ListBinary(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountBinary(ctx, parameter)
}
