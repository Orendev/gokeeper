package user

import (
	domainUser "github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/protobuff"
)

func ToUserResponse(response *domainUser.User) *protobuff.CreateUserResponse {
	return &protobuff.CreateUserResponse{
		ID:         response.ID().String(),
		Email:      response.Email().String(),
		Name:       response.Name().String(),
		Surname:    response.Surname().String(),
		Patronymic: response.Patronymic().String(),
	}
}
