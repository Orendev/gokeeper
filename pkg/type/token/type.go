package token

type Token string

func (t Token) String() string {
	return string(t)
}

func New(token string) *Token {

	t := Token(token)

	return &t
}
