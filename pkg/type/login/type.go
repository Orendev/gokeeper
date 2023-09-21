package login

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("login must be less than or equal to %d characters", MaxLength)
)

type Login string

// New - constructor Login
func New(login string) (*Login, error) {
	if len([]rune(login)) > MaxLength {
		return nil, ErrWrongLength
	}

	l := Login(login)

	return &l, nil
}

func (l Login) String() string {
	return string(l)
}

func (l Login) IsEmpty() bool {
	return len(strings.TrimSpace(l.String())) == 0
}
