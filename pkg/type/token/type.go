package token

type Token string

func (t Token) String() string {
	return string(t)
}

func New(token string) (*Token, error) {

	t := Token(token)

	return &t, nil
}
