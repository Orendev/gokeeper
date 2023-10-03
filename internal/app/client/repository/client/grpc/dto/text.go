package dto

import (
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/mashingan/smapping"
)

type Data struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    string `json:"user_id"`
	Data      []byte `json:"data"`
	Comment   []byte `json:"comment"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListData struct {
	Total  uint64 `json:"total"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
	Data   []Data
}

// FromListTextResponseToDto converts json body request to a ListTextResponse struct
func FromListTextResponseToDto(source *protobuff.ListTextResponse) (*ListData, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToListDataDto(mapped)
}

func fromMappedToListDataDto(mapped smapping.Mapped) (*ListData, error) {
	listData := ListData{}
	err := smapping.FillStruct(&listData, mapped)
	if err != nil {
		return nil, err
	}

	return &listData, nil
}
