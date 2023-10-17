package binary

import (
	"time"

	domainBinary "github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	"github.com/Orendev/gokeeper/pkg/type/title"
)

type CreateBinaryArgs struct {
	Title   string `json:"title"`
	Data    string `json:"data"`
	Comment string `json:"comment"`
	UserID  string `json:"user_id"`
}

type UpdateBinaryArgs struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Data    string `json:"data"`
	Comment string `json:"comment"`
	UserID  string `json:"user_id"`
}

type DeleteBinaryArgs struct {
	ID string `json:"id"`
}

type ListBinaryArgs struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

func ToEncCreateBinary(enc *encryption.Enc, args *CreateBinaryArgs) (*domainBinary.BinaryData, error) {

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

	return domainBinary.New(
		converter.StringToUUID(args.UserID),
		*titleObj,
		dataEnc,
		commentEnc,
	)
}

func ToEncUpdateBinary(enc *encryption.Enc, args *UpdateBinaryArgs) (*domainBinary.BinaryData, error) {

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

	return domainBinary.NewWithID(
		converter.StringToUUID(args.ID),
		converter.StringToUUID(args.UserID),
		*titleObj,
		dataEnc,
		commentEnc,
		time.Now().UTC(),
		time.Now().UTC(),
	)
}

func ToDecBinary(enc *encryption.Enc, val *domainBinary.BinaryData) (*domainBinary.BinaryData, error) {

	dataDec, err := enc.DecryptByte(val.Data())
	if err != nil {
		return nil, err
	}

	commentDec, err := enc.DecryptByte(val.Comment())
	if err != nil {
		return nil, err
	}

	return domainBinary.NewWithID(
		val.ID(),
		val.UserID(),
		val.Title(),
		dataDec,
		commentDec,
		val.CreatedAt(),
		val.UpdatedAt(),
	)
}
