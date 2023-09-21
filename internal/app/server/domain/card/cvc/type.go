package cvc

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	MaxLength      uint64 = 999
	ErrWrongLength        = errors.Errorf("Age must be less than or equal to %d", MaxLength)
)

type CVC uint8

func (c CVC) String() string {
	return strconv.FormatUint(uint64(c), 10)
}

func New(cvc uint64) (*CVC, error) {
	if cvc > MaxLength {
		return nil, ErrWrongLength
	}
	c := CVC(cvc)

	return &c, nil
}

func (c CVC) IsEmpty() bool {
	return len(strings.TrimSpace(c.String())) == 0
}
