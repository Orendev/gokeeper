package dto

import (
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/mashingan/smapping"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// FromLoginUserResponseToDto converts json body request to a LoginUserResponse struct
func FromLoginUserResponseToDto(source *protobuff.LoginUserResponse) (*User, error) {

	mapped := smapping.MapFields(source)

	return fromMappedToUserDto(mapped)
}

// FromRegisterUserResponseToDto converts json body request to a RegisterUserResponse struct
func FromRegisterUserResponseToDto(source *protobuff.RegisterUserResponse) (*User, error) {
	mapped := smapping.MapFields(source)

	return fromMappedToUserDto(mapped)
}

func fromMappedToUserDto(mapped smapping.Mapped) (*User, error) {
	user := User{}
	err := smapping.FillStruct(&user, mapped)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
