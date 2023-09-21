package storage

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
)

// User Interface for interacting with the use case repository.
type User interface {
	CreateUser(users ...*user.User) ([]*user.User, error)

	UserReader
}

// UserReader user-readable interface
type UserReader interface {
	FindUser(email email.Email) (*user.User, error)
}
