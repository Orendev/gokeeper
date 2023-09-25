package role

import (
	"github.com/pkg/errors"
)

type Role int

const (
	Admin Role = iota
	User
)

var (
	ErrWrongRole = errors.Errorf("specify the correct role")
)

func New(role string) (*Role, error) {
	var access = 1

	switch role {
	case "admin":
		access = 0
	case "user":
		access = 1
	default:
		return nil, ErrWrongRole
	}

	r := Role(access)

	return &r, nil
}

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case User:
		return "user"
	}

	return "user"
}
