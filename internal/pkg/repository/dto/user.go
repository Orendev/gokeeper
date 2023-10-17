package dto

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/role"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/google/uuid"
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
	u := User{}
	err := smapping.FillStruct(&u, mapped)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func ToDomainUser(dto User) (*user.User, error) {

	nameObject, err := name.New(dto.Name)
	if err != nil {
		return nil, err
	}

	emailObject, err := email.New(dto.Email)
	if err != nil {
		return nil, err
	}

	passwordObject, err := password.New(dto.Password)
	if err != nil {
		return nil, err
	}

	roleObject, err := role.New(dto.Role)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, dto.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, dto.UpdatedAt)
	if err != nil {
		return nil, err
	}

	tokenObject := token.New(dto.AccessToken)
	if err != nil {
		return nil, err
	}

	result, err := user.NewWithID(
		id,
		*passwordObject,
		*emailObject,
		*nameObject,
		*roleObject,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}

	result.SetToken(*tokenObject)

	return result, nil
}
