package user

import (
	"time"

	domainUser "github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/protobuff"
)

func ToRegisterUserResponse(response *domainUser.User) *protobuff.RegisterUserResponse {
	return &protobuff.RegisterUserResponse{
		ID:          response.ID().String(),
		Email:       response.Email().String(),
		Name:        response.Name().String(),
		Role:        response.Role().String(),
		AccessToken: response.Token().String(),
		CreatedAt:   response.CreatedAt().Format(time.RFC3339),
		UpdatedAt:   response.UpdatedAt().Format(time.RFC3339),
	}
}

func ToLoginUserResponse(response *domainUser.User) *protobuff.LoginUserResponse {
	return &protobuff.LoginUserResponse{
		ID:          response.ID().String(),
		Email:       response.Email().String(),
		Name:        response.Name().String(),
		Role:        response.Role().String(),
		AccessToken: response.Token().String(),
		CreatedAt:   response.CreatedAt().Format(time.RFC3339),
		UpdatedAt:   response.UpdatedAt().Format(time.RFC3339),
	}
}
