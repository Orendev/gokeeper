package name

import "github.com/pkg/errors"

var (
	MaxLength         = 50
	MinLength         = 3
	ErrWrongMaxLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)
	ErrWrongMinLength = errors.Errorf("name must be greater than or equal to %d characters", MinLength)
)

type Name string

func (n Name) String() string {
	return string(n)
}

func New(name string) (*Name, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongMaxLength
	}
	if len([]rune(name)) < MinLength {
		return nil, ErrWrongMinLength
	}
	n := Name(name)

	return &n, nil
}
