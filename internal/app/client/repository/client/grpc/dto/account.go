package dto

import (
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/mashingan/smapping"
)

type DeleteAccount struct {
	ID string `json:"id"`
}

type CreateAccount struct {
	ID string `json:"id"`
}

type UpdateAccount struct {
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
	Login     string `json:"login"`
	Password  string `json:"password"`
	URL       string `json:"url"`
	Comment   string `json:"comment"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// FromCreateAccountResponseToDto converts json body request to a CreateAccountResponse struct
func FromCreateAccountResponseToDto(source *protobuff.CreateAccountResponse) (*CreateAccount, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToCreateAccountDto(mapped)
}

func fromMappedToCreateAccountDto(mapped smapping.Mapped) (*CreateAccount, error) {
	createAccount := CreateAccount{}
	err := smapping.FillStruct(&createAccount, mapped)
	if err != nil {
		return nil, err
	}

	return &createAccount, nil
}

// FromUpdateAccountResponseToDto converts json body request to a CreateAccountResponse struct
func FromUpdateAccountResponseToDto(source *protobuff.UpdateAccountResponse) (*UpdateAccount, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToUpdateAccountDto(mapped)
}

func fromMappedToUpdateAccountDto(mapped smapping.Mapped) (*UpdateAccount, error) {
	updateAccount := UpdateAccount{}
	err := smapping.FillStruct(&updateAccount, mapped)
	if err != nil {
		return nil, err
	}

	return &updateAccount, nil
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

func fromMappedToDeleteAccountDto(mapped smapping.Mapped) (*DeleteAccount, error) {
	deleteAccount := DeleteAccount{}
	err := smapping.FillStruct(&deleteAccount, mapped)
	if err != nil {
		return nil, err
	}

	return &deleteAccount, nil
}

// FromDeleteAccountResponseToDto converts json body request to a CreateAccountResponse struct
func FromDeleteAccountResponseToDto(source *protobuff.DeleteAccountResponse) (*DeleteAccount, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToDeleteAccountDto(mapped)
}

func fromMappedToAccountDto(mapped smapping.Mapped) (*Account, error) {
	account := Account{}
	err := smapping.FillStruct(&account, mapped)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
