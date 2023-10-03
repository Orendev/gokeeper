package converter

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
)

func ToDomainText(dto dto.Data) (*text.TextData, error) {

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

	result, err := text.NewWithID(
		id,
		userID,
		*titleObj,
		dto.Data,
		dto.Comment,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ToDomainTexts(dto dto.ListData) ([]*text.TextData, error) {

	result := make([]*text.TextData, len(dto.Data))

	for i, val := range dto.Data {
		a, err := ToDomainText(val)
		if err != nil {
			return nil, err
		}

		result[i] = a
	}

	return result, nil
}
