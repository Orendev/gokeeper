package converter

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/google/uuid"
)

func ToDomainCard(dto dto.Card) (*card.CardData, error) {

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

	return card.NewWithID(
		id,
		userID,
		dto.CardNumber,
		dto.CardName,
		dto.CVC,
		dto.CardDate,
		dto.Comment,
		createdAt,
		updatedAt,
	)
}

func ToDomainCards(dto dto.ListCard) ([]*card.CardData, error) {

	result := make([]*card.CardData, len(dto.Data))

	for i, val := range dto.Data {
		a, err := ToDomainCard(val)
		if err != nil {
			return nil, err
		}

		result[i] = a
	}

	return result, nil
}
