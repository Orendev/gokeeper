package dto

import (
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/mashingan/smapping"
)

type CreateAccount struct {
	ID string `json:"id"`
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

func fromMappedToAccountDto(mapped smapping.Mapped) (*Account, error) {
	account := Account{}
	err := smapping.FillStruct(&account, mapped)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
