package password

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("password must be less than or equal to %d characters", MaxLength)
)

type Password string

func New(password string) (*Password, error) {

	if len([]rune(password)) > MaxLength {
		return nil, ErrWrongLength
	}

	p := Password(password)

	return &p, nil
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Byte() []byte {
	return []byte(p)
}

func (p Password) IsEmpty() bool {
	return len(strings.TrimSpace(p.String())) == 0
}

func (p Password) Equal(password Password) bool {
	return p == password
}
