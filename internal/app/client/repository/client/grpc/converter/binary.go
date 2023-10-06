package converter

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
)

func ToDomainBinary(dto dto.Data) (*binary.BinaryData, error) {

	titleObj, err := title.New(dto.Title)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(dto.UserID)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, dto.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, dto.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return binary.NewWithID(
		id,
		userID,
		*titleObj,
		dto.Data,
		dto.Comment,
		createdAt,
		updatedAt,
	)
}

func ToDomainBinaries(dto dto.ListData) ([]*binary.BinaryData, error) {

	result := make([]*binary.BinaryData, len(dto.Data))

	for i, val := range dto.Data {
		a, err := ToDomainBinary(val)
		if err != nil {
			return nil, err
		}

		result[i] = a
	}

	return result, nil
}
