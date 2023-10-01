package account

import (
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/pkg/errors"
)

type CreateAccountArgs struct {
	Title    string `json:"title"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

type UpdateAccountArgs struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

type DeleteAccountArgs struct {
	ID string `json:"id"`
}

type ListAccountArgs struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

type encryptable interface {
	CreateAccountArgs | UpdateAccountArgs
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

func ToEncAccountArgs[T encryptable](enc *encryption.Enc, args *T) error {
	var err error
	switch v := interface{}(args).(type) {
	case *UpdateAccountArgs:
		loginEnc, err := enc.Encrypt(v.Login)
		if err != nil {
			return err
		}
		v.Login = loginEnc

		passwordEnc, err := enc.Encrypt(v.Password)
		if err != nil {
			return err
		}
		v.Password = passwordEnc

		commentEnc, err := enc.Encrypt(v.Comment)
		if err != nil {
			return err
		}
		v.Comment = commentEnc

	case *CreateAccountArgs:
		loginEnc, err := enc.Encrypt(v.Login)
		if err != nil {
			return err
		}
		v.Login = loginEnc

		passwordEnc, err := enc.Encrypt(v.Password)
		if err != nil {
			return err
		}
		v.Password = passwordEnc

		commentEnc, err := enc.Encrypt(v.Comment)
		if err != nil {
			return err
		}
		v.Comment = commentEnc
	default:
		return errors.Errorf("I don't know about type %T!\n", v)
	}

	return err
}

func ToDecAccount(enc *encryption.Enc, val *account.Account) (*account.Account, error) {

	loginDec, err := enc.Decrypt(val.Login().String())
	if err != nil {
		return nil, err
	}

	loginObj, err := login.New(loginDec)
	if err != nil {
		return nil, err
	}

	passwordDec, err := enc.Decrypt(val.Password().String())
	if err != nil {
		return nil, err
	}

	passwordObj, err := password.New(passwordDec)
	if err != nil {
		return nil, err
	}

	commentDec, err := enc.Decrypt(val.Comment().String())
	if err != nil {
		return nil, err
	}

	commentObj, err := comment.New(commentDec)
	if err != nil {
		return nil, err
	}

	dAccount, err := account.NewWithID(
		val.ID(),
		val.Title(),
		*loginObj,
		*passwordObj,
		val.URL(),
		*commentObj,
		val.IsDeleted(),
		val.CreatedAt(),
		val.UpdatedAt(),
	)
	if err != nil {
		return nil, err
	}

	return dAccount, nil
}
