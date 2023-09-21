package useCase

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
)

// User Interface for interaction between delivery and use case
type User interface {
	Create(users ...*user.User) ([]*user.User, error)

	UserReader
}

type UserReader interface {
	Find(email email.Email) (*user.User, error)
}
