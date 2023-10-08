package card

import (
	"time"

	domainCard "github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/google/uuid"
)

func ToCreateCardResponse(account *domainCard.CardData) *protobuff.CreateCardResponse {
	return &protobuff.CreateCardResponse{
		ID: account.ID().String(),
	}
}

func ToUpdateCardResponse(account *domainCard.CardData) *protobuff.UpdateCardResponse {
	return &protobuff.UpdateCardResponse{
		ID: account.ID().String(),
	}
}

func ToDeleteCardResponse(id uuid.UUID) *protobuff.DeleteCardResponse {
	return &protobuff.DeleteCardResponse{
		ID: id.String(),
	}
}

func ToListCardResponse(list *domainCard.ListCardViewModel) *protobuff.ListCardResponse {
	data := []*protobuff.Card{}

	for _, value := range list.Data {
		data = append(data, &protobuff.Card{
			ID:         value.ID().String(),
			CardName:   value.CardName(),
			CardNumber: value.CardNumber(),
			CardDate:   value.CardDate(),
			CVC:        value.CVC(),
			Comment:    value.Comment(),
			CreatedAt:  value.CreatedAt().Format(time.RFC3339),
			UpdatedAt:  value.UpdatedAt().Format(time.RFC3339),
			UserID:     value.UserID().String(),
		})
	}

	return &protobuff.ListCardResponse{
		Total:  list.Total,
		Limit:  list.Limit,
		Offset: list.Offset,
		Data:   data,
	}
}
