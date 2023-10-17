package dto

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
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

func fromMappedToTextDto(mapped smapping.Mapped) (*Data, error) {
	d := Data{}
	err := smapping.FillStruct(&d, mapped)
	if err != nil {
		return nil, err
	}

	return &d, nil
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

func ToDomainText(dto Data) (*text.TextData, error) {

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

	return text.NewWithID(
		id,
		userID,
		*titleObj,
		dto.Data,
		dto.Comment,
		createdAt,
		updatedAt,
	)
}

func ToDomainTexts(dto ListData) ([]*text.TextData, error) {

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

// BINARY

// FromCreateBinaryResponseToDto converts json body request to a CreateTextResponse struct
func FromCreateBinaryResponseToDto(source *protobuff.CreateBinaryResponse) (*Data, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToTextDto(mapped)
}

// FromUpdateBinaryResponseToDto converts json body request to a UpdateTextResponse struct
func FromUpdateBinaryResponseToDto(source *protobuff.UpdateBinaryResponse) (*Data, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToTextDto(mapped)
}

// FromListBinaryResponseToDto converts json body request to a ListTextResponse struct
func FromListBinaryResponseToDto(source *protobuff.ListBinaryResponse) (*ListData, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToListDataDto(mapped)
}

func ToDomainBinary(dto Data) (*binary.BinaryData, error) {

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

func ToDomainBinaries(dto ListData) ([]*binary.BinaryData, error) {

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
