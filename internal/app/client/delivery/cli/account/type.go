package account

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	"github.com/Orendev/gokeeper/pkg/type/title"
)

// CreateAccountArgs data structure for creating an account
type CreateAccountArgs struct {
	Title    string `json:"title"`
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

// UpdateAccountArgs data structure for update an account
type UpdateAccountArgs struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Title    string `json:"title"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

// DeleteAccountArgs data structure for delete an account
type DeleteAccountArgs struct {
	ID string `json:"id"`
}

// ListAccountArgs data structure for getting a list of accounts
type ListAccountArgs struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

// ToEncCreateAccount data encoding account creation
func ToEncCreateAccount(enc *encryption.Enc, args *CreateAccountArgs) (*account.Account, error) {
	var err error

	titleObj, err := title.New(args.Title)
	if err != nil {
		return nil, err
	}

	loginEnc, err := enc.EncryptByte([]byte(args.Login))
	if err != nil {
		return nil, err
	}

	passwordEnc, err := enc.EncryptByte([]byte(args.Password))
	if err != nil {
		return nil, err
	}

	urlEnc, err := enc.EncryptByte([]byte(args.URL))
	if err != nil {
		return nil, err
	}

	commentEnc, err := enc.EncryptByte([]byte(args.Comment))
	if err != nil {
		return nil, err
	}

	return account.New(
		converter.StringToUUID(args.UserID),
		*titleObj,
		loginEnc,
		passwordEnc,
		urlEnc,
		commentEnc,
	)
}

// ToEncUpdateAccount data encoding account update
func ToEncUpdateAccount(enc *encryption.Enc, args *UpdateAccountArgs) (*account.Account, error) {
	var err error

	titleObj, err := title.New(args.Title)
	if err != nil {
		return nil, err
	}

	loginEnc, err := enc.EncryptByte([]byte(args.Login))
	if err != nil {
		return nil, err
	}

	passwordEnc, err := enc.EncryptByte([]byte(args.Password))
	if err != nil {
		return nil, err
	}

	urlEnc, err := enc.EncryptByte([]byte(args.URL))
	if err != nil {
		return nil, err
	}

	commentEnc, err := enc.EncryptByte([]byte(args.Comment))
	if err != nil {
		return nil, err
	}

	return account.NewWithID(
		converter.StringToUUID(args.ID),
		converter.StringToUUID(args.UserID),
		*titleObj,
		loginEnc,
		passwordEnc,
		urlEnc,
		commentEnc,
		time.Now().UTC(),
		time.Now().UTC(),
	)
}

// ToDecAccount decoding account data
func ToDecAccount(enc *encryption.Enc, val *account.Account) (*account.Account, error) {

	loginDec, err := enc.DecryptByte(val.Login())
	if err != nil {
		return nil, err
	}

	passwordDec, err := enc.DecryptByte(val.Password())
	if err != nil {
		return nil, err
	}

	urlDec, err := enc.DecryptByte(val.URL())
	if err != nil {
		return nil, err
	}

	commentDec, err := enc.DecryptByte(val.Comment())
	if err != nil {
		return nil, err
	}

	return account.NewWithID(
		val.ID(),
		val.UserID(),
		val.Title(),
		loginDec,
		passwordDec,
		urlDec,
		commentDec,
		val.CreatedAt(),
		val.UpdatedAt(),
	)
}
