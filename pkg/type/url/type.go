package url

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	regexpWebAddress  = regexp.MustCompile(`^((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)$`)
	ErrWrongFormatURL = errors.New("invalid url format")
)

type URL struct {
	value string
}

func (u URL) String() string {
	return u.value
}

func New(url string) (*URL, error) {
	if len(url) == 0 {
		return &URL{}, nil
	}

	if !regexpWebAddress.MatchString(strings.ToLower(url)) {
		return nil, ErrWrongFormatURL
	}

	return &URL{
		value: url,
	}, nil
}
