package title

import "github.com/pkg/errors"

var (
	MaxLength      = 255
	ErrWrongLength = errors.Errorf("title must be less than or equal to %d characters", MaxLength)
)

type Title string

func (t Title) String() string {
	return string(t)
}

func New(title string) (*Title, error) {
	if len([]rune(title)) > MaxLength {
		return nil, ErrWrongLength
	}

	t := Title(title)

	return &t, nil
}
