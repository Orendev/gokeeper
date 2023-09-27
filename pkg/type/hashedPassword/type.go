package hashedPassword

import (
	"fmt"
	"strings"

	"github.com/Orendev/gokeeper/pkg/type/password"
	"golang.org/x/crypto/bcrypt"
)

type HashedPassword string

func New(password string) (*HashedPassword, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	encryptedPassword := string(hashedPassword)
	p := HashedPassword(encryptedPassword)

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
