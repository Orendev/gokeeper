package comment

import "github.com/pkg/errors"

var (
	MaxLength      = 255
	ErrWrongLength = errors.Errorf("comment must be less than or equal to %d characters", MaxLength)
)

type Comment string

func (c Comment) String() string {
	return string(c)
}

func New(comment string) (*Comment, error) {
	if len([]rune(comment)) > MaxLength {
		return nil, ErrWrongLength
	}

	c := Comment(comment)

	return &c, nil
}
