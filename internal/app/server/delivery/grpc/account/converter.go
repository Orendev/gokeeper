package account

import (
	"time"

	domainAccount "github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/protobuff"
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

func ToListAccountResponse(list *domainAccount.ListAccountViewModel) *protobuff.ListAccountResponse {
	data := []*protobuff.Account{}
	for _, value := range list.Data {
		data = append(data, &protobuff.Account{
			ID:        value.ID().String(),
			Title:     value.Title().String(),
			Login:     value.Login(),
			Password:  value.Password(),
			URL:       value.URL(),
			Comment:   value.Comment(),
			CreatedAt: value.CreatedAt().Format(time.RFC3339),
			UpdatedAt: value.UpdatedAt().Format(time.RFC3339),
			UserID:    value.UserID().String(),
		})
	}

	return &protobuff.ListAccountResponse{
		Total:  list.Total,
		Limit:  list.Limit,
		Offset: list.Offset,
		Data:   data,
	}
}
