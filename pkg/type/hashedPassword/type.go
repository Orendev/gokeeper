package hashedPassword

import (
	"fmt"
	"strings"

	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	MaxLength          = 50
	ErrPasswordTooLong = errors.Errorf("password length exceeds %d characters", MaxLength)
)

type HashedPassword string

func New(password string) (*HashedPassword, error) {

	if len([]rune(password)) > MaxLength {
		return nil, ErrPasswordTooLong
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	p := HashedPassword(hashedPassword)

	return &p, nil
}

func (p HashedPassword) String() string {
	return string(p)
}

func (p HashedPassword) Byte() []byte {
	return []byte(p)
}

func (p HashedPassword) IsEmpty() bool {
	return len(strings.TrimSpace(p.String())) == 0
}

// CompareHashAndPassword compares a bcrypt hashed password with its possible
func (p HashedPassword) CompareHashAndPassword(password password.Password) error {
	return bcrypt.CompareHashAndPassword(p.Byte(), password.Byte())
}