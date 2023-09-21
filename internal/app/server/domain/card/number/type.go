package number

import "strings"

type Number string

func (n Number) String() string {
	return string(n)
}

func New(number string) (*Number, error) {

	n := Number(number)

	return &n, nil
}

func (n Number) IsEmpty() bool {
	return len(strings.TrimSpace(n.String())) == 0
}
