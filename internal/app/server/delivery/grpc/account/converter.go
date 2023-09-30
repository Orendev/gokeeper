package account

import (
	domainAccount "github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
)

func ToCreateAccountResponse(account *domainAccount.Account) *protobuff.CreateAccountResponse {
	return &protobuff.CreateAccountResponse{
		ID: account.ID().String(),
	}
}
