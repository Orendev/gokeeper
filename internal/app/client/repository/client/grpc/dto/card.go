package dto

import (
	"github.com/Orendev/gokeeper/pkg/protobuff"
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
