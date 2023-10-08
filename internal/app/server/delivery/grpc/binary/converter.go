package binary

import (
	"time"

	domainBinary "github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/google/uuid"
)

func ToCreateBinaryResponse(account *domainBinary.BinaryData) *protobuff.CreateBinaryResponse {
	return &protobuff.CreateBinaryResponse{
		ID: account.ID().String(),
	}
}

func ToUpdateBinaryResponse(account *domainBinary.BinaryData) *protobuff.UpdateBinaryResponse {
	return &protobuff.UpdateBinaryResponse{
		ID: account.ID().String(),
	}
}

func ToDeleteBinaryResponse(id uuid.UUID) *protobuff.DeleteBinaryResponse {
	return &protobuff.DeleteBinaryResponse{
		ID: id.String(),
	}
}

func ToListBinaryResponse(list *domainBinary.ListBinaryViewModel) *protobuff.ListBinaryResponse {
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

	return &protobuff.ListBinaryResponse{
		Total:  list.Total,
		Limit:  list.Limit,
		Offset: list.Offset,
		Data:   data,
	}
}
