package password

import (
	"strings"
)

type Password string

func New(password string) (*Password, error) {

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
