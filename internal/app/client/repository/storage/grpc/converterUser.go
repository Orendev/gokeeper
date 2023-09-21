package grpc

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user/name"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user/role"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/google/uuid"
)

func (r Repository) toDomainUser(res *protobuff.CreateUserResponse, userObject user.User) (*user.User, error) {

	emailObject, err := email.New(res.GetEmail())
	if err != nil {
		return nil, err
	}

	roleObject, err := role.New("user")
	if err != nil {
		return nil, err
	}

	nameObject, err := name.New(res.GetName())
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(res.GetID())
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.DateTime, res.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.DateTime, res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	result, err := user.NewWithID(
		id,
		userObject.Password(),
		*emailObject,
		*roleObject,
		*nameObject,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
