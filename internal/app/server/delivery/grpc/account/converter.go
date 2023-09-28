package account

import (
	domainAccount "github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"time"
)

func ToCreateAccountResponse(account *domainAccount.Account) *protobuff.CreateAccountResponse {
	return &protobuff.CreateAccountResponse{
		ID:        account.ID().String(),
		Title:     account.Title().String(),
		Login:     account.Login().String(),
		Password:  account.Password().String(),
		URL:       account.URL().String(),
		CreatedAt: account.CreatedAt().Format(time.RFC3339),
		UpdatedAt: account.UpdatedAt().Format(time.RFC3339),
	}
}
