package dto

import (
	"fmt"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type DeleteAccount struct {
	ID string `json:"id"`
}

type ListAccount struct {
	Total  uint64 `json:"total"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
	Data   []Account
}

type Account struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    string `json:"user_id"`
	Login     []byte `json:"login"`
	Password  []byte `json:"password"`
	URL       []byte `json:"url"`
	Comment   []byte `json:"comment"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// FromListAccountResponseToDto converts json body request to a ListAccountResponse struct
func FromListAccountResponseToDto(source *protobuff.ListAccountResponse) (*ListAccount, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToListAccountDto(mapped)
}

func fromMappedToListAccountDto(mapped smapping.Mapped) (*ListAccount, error) {
	listAccount := ListAccount{}
	err := smapping.FillStruct(&listAccount, mapped)
	if err != nil {
		return nil, err
	}

	return &listAccount, nil
}

func ToDomainAccount(dto Account) (*account.Account, error) {

	titleObj, err := title.New(dto.Title)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, dto.CreatedAt)
	if err != nil {
		fmt.Println("fff", dto)
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, dto.UpdatedAt)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(dto.UserID)
	if err != nil {
		return nil, err
	}

	return account.NewWithID(
		id,
		userID,
		*titleObj,
		dto.Login,
		dto.Password,
		dto.URL,
		dto.Comment,
		createdAt,
		updatedAt,
	)
}

func ToDomainAccounts(dto ListAccount) ([]*account.Account, error) {

	result := make([]*account.Account, len(dto.Data))

	for i, val := range dto.Data {
		a, err := ToDomainAccount(val)
		if err != nil {
			return nil, err
		}

		result[i] = a
	}

	return result, nil
}
