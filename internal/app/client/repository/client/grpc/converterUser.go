package grpc

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/token"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/role"
	"github.com/google/uuid"
)

func toDomainUser(dto dto.User) (*user.User, error) {

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

	tokenObject, err := token.New(dto.AccessToken)
	if err != nil {
		return nil, err
	}

	result, err := user.NewWithID(
		id,
		*passwordObject,
		*emailObject,
		*roleObject,
		*nameObject,
		*tokenObject,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
