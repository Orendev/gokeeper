package binary

import (
	"context"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {
	return uc.adapterStorage.CreateBinary(ctx, binary)
}

func (uc *UseCase) Update(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {
	return uc.adapterStorage.UpdateBinary(ctx, binary)
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteBinary(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) (*binary.ListBinaryViewModel, error) {
	return uc.adapterStorage.ListBinary(ctx, parameter)
}

func (uc *UseCase) Count(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	return uc.adapterStorage.CountBinary(ctx, parameter)
}
