package user

import (
	domainUser "github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/protobuff"
)

func ToUserResponse(response *domainUser.User) *protobuff.RegisterUserResponse {
	return &protobuff.RegisterUserResponse{
		ID:          response.ID().String(),
		Email:       response.Email().String(),
		Name:        response.Name().String(),
		Role:        response.Role().String(),
		AccessToken: response.Token().String(),
		CreatedAt:   response.CreatedAt().String(),
		UpdatedAt:   response.UpdatedAt().String(),
	}
}
