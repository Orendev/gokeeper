package account

import (
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
)

type CreateAccountArgs struct {
	Title    string `json:"title"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

func ToEncCreateAccountArgs(enc *encryption.Enc, args *CreateAccountArgs) error {
	var err error
	loginEnc, err := enc.Encrypt(args.Login)
	if err != nil {
		return err
	}
	args.Login = loginEnc

	passwordEnc, err := enc.Encrypt(args.Password)
	if err != nil {
		return err
	}
	args.Password = passwordEnc

	commentEnc, err := enc.Encrypt(args.Comment)
	if err != nil {
		return err
	}
	args.Comment = commentEnc

	return err
}
