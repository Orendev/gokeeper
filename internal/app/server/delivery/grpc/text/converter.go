package text

import (
	"time"

	domainText "github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/google/uuid"
)

func ToCreateTextResponse(account *domainText.TextData) *protobuff.CreateTextResponse {
	return &protobuff.CreateTextResponse{
		ID: account.ID().String(),
	}
}

func ToUpdateTextResponse(account *domainText.TextData) *protobuff.UpdateTextResponse {
	return &protobuff.UpdateTextResponse{
		ID: account.ID().String(),
	}
}

func ToDeleteTextResponse(id uuid.UUID) *protobuff.DeleteTextResponse {
	return &protobuff.DeleteTextResponse{
		ID: id.String(),
	}
}

func ToListTextResponse(list *domainText.ListTextViewModel) *protobuff.ListTextResponse {
	data := []*protobuff.Data{}

	for _, value := range list.Data {
		data = append(data, &protobuff.Data{
			ID:        value.ID().String(),
			Title:     value.Title().String(),
			Data:      value.Data(),
			Comment:   value.Comment(),
			CreatedAt: value.CreatedAt().Format(time.RFC3339),
			UpdatedAt: value.UpdatedAt().Format(time.RFC3339),
			UserID:    value.UserID().String(),
		})
	}

	return &protobuff.ListTextResponse{
		Total:  list.Total,
		Limit:  list.Limit,
		Offset: list.Offset,
		Data:   data,
	}
}
