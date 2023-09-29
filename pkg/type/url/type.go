package url

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

	return &URL{
		value: url,
	}, nil
}
