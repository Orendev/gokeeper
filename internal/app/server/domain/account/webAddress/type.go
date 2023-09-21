package webAddress

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	regexpWebAddress         = regexp.MustCompile(`^((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)$`)
	ErrWrongFormatWebAddress = errors.New("invalid email format")
)

type WebAddress struct {
	value string
}

func (w WebAddress) String() string {
	return w.value
}

func New(webAddress string) (*WebAddress, error) {
	if len(webAddress) == 0 {
		return &WebAddress{}, nil
	}

	if !regexpWebAddress.MatchString(strings.ToLower(webAddress)) {
		return nil, ErrWrongFormatWebAddress
	}

	return &WebAddress{
		value: webAddress,
	}, nil
}
