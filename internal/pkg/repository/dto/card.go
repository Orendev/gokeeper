package dto

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type Card struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	UserID     string `json:"user_id"`
	CardNumber []byte `json:"card_number"`
	CardName   []byte `json:"card_name"`
	CardDate   []byte `json:"card_date"`
	CVC        []byte `json:"cvc"`
	Comment    []byte `json:"comment"`
	IsDelete   bool   `json:"is_deleted"`
}

type ListCard struct {
	Total  uint64 `json:"total"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
	Data   []Card
}

// FromListCardResponseToDto converts json body request to a ListCardResponse struct
func FromListCardResponseToDto(source *protobuff.ListCardResponse) (*ListCard, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToListCardDto(mapped)
}

func fromMappedToListCardDto(mapped smapping.Mapped) (*ListCard, error) {
	listData := ListCard{}
	err := smapping.FillStruct(&listData, mapped)
	if err != nil {
		return nil, err
	}

	return &listData, nil
}

func ToDomainCard(dto Card) (*card.CardData, error) {

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

func ToDomainCards(dto ListCard) ([]*card.CardData, error) {

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
