package body

type Body struct {
	value string
}

// New - constructor Body
func New(body string) (*Body, error) {

	return &Body{body}, nil
}

func (b Body) Value() string {
	return b.value
}
