package body

type Body struct {
	value []byte
}

// New - constructor Body
func New(body []byte) (*Body, error) {

	return &Body{body}, nil
}

func (b Body) Value() []byte {
	return b.value
}
