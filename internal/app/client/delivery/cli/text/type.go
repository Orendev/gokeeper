package text

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	"github.com/Orendev/gokeeper/pkg/type/title"
)

type CreateTextArgs struct {
	Title   string `json:"title"`
	Data    string `json:"data"`
	Comment string `json:"comment"`
	UserID  string `json:"user_id"`
}

type UpdateTextArgs struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Data    string `json:"data"`
	Comment string `json:"comment"`
	UserID  string `json:"user_id"`
}

type DeleteTextArgs struct {
	ID string `json:"id"`
}

type ListTextArgs struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

func ToEncCreateText(enc *encryption.Enc, args *CreateTextArgs) (*text.TextData, error) {

	titleObj, err := title.New(args.Title)
	if err != nil {
		return nil, err
	}

	dataEnc, err := enc.EncryptByte([]byte(args.Data))
	if err != nil {
		return nil, err
	}

	commentEnc, err := enc.EncryptByte([]byte(args.Comment))
	if err != nil {
		return nil, err
	}

	return text.New(
		converter.StringToUUID(args.UserID),
		*titleObj,
		dataEnc,
		commentEnc,
	)
}

func ToEncUpdateText(enc *encryption.Enc, args *UpdateTextArgs) (*text.TextData, error) {

	titleObj, err := title.New(args.Title)
	if err != nil {
		return nil, err
	}

	dataEnc, err := enc.EncryptByte([]byte(args.Data))
	if err != nil {
		return nil, err
	}

	commentEnc, err := enc.EncryptByte([]byte(args.Comment))
	if err != nil {
		return nil, err
	}

	return text.NewWithID(
		converter.StringToUUID(args.ID),
		converter.StringToUUID(args.UserID),
		*titleObj,
		dataEnc,
		commentEnc,
		time.Now().UTC(),
		time.Now().UTC(),
	)
}

func ToDecText(enc *encryption.Enc, val *text.TextData) (*text.TextData, error) {

	dataDec, err := enc.DecryptByte(val.Data())
	if err != nil {
		return nil, err
	}

	commentDec, err := enc.DecryptByte(val.Comment())
	if err != nil {
		return nil, err
	}

	return text.NewWithID(
		val.ID(),
		val.UserID(),
		val.Title(),
		dataDec,
		commentDec,
		val.CreatedAt(),
		val.UpdatedAt(),
	)
}
