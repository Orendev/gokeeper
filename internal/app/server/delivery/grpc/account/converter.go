package account

import (
	"time"

	domainAccount "github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func ToCreateAccountResponse(account *domainAccount.Account) *protobuff.CreateAccountResponse {
	return &protobuff.CreateAccountResponse{
		ID: account.ID().String(),
	}
}

func ToUpdateAccountResponse(account *domainAccount.Account) *protobuff.UpdateAccountResponse {
	return &protobuff.UpdateAccountResponse{
		ID: account.ID().String(),
	}
}

func ToDeleteAccountResponse(id uuid.UUID) *protobuff.DeleteAccountResponse {
	return &protobuff.DeleteAccountResponse{
		ID: id.String(),
	}
}

func ToListAccountResponse(accounts []*domainAccount.Account, parameter queryParameter.QueryParameter, total uint64) *protobuff.ListAccountResponse {
	data := []*protobuff.AccountResponse{}
	for _, value := range accounts {
		data = append(data, &protobuff.AccountResponse{
			ID:        value.ID().String(),
			Title:     value.Title().String(),
			Login:     value.Login().String(),
			Password:  value.Password().String(),
			URL:       value.URL().String(),
			Comment:   value.Comment().String(),
			CreatedAt: value.CreatedAt().Format(time.RFC3339),
			UpdatedAt: value.UpdatedAt().Format(time.RFC3339),
		})
	}

	return &protobuff.ListAccountResponse{
		Total:  total,
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Data:   data,
	}
}
