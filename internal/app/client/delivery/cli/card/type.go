package card

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/tools/converter"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
)

type CreateCardArgs struct {
	ID         string `json:"id"`
	CardNumber string `json:"card_number"`
	CardName   string `json:"card_name"`
	CardDate   string `json:"card_date"`
	CVC        string `json:"cvc"`
	Comment    string `json:"comment"`
	UserID     string `json:"user_id"`
}

type UpdateCardArgs struct {
	ID         string `json:"id"`
	CardNumber string `json:"card_number"`
	CardName   string `json:"card_name"`
	CardDate   string `json:"card_date"`
	CVC        string `json:"cvc"`
	Comment    string `json:"comment"`
	UserID     string `json:"user_id"`
}

type DeleteCardArgs struct {
	ID string `json:"id"`
}

type ListCardArgs struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

func ToEncCreateCard(enc *encryption.Enc, args *CreateCardArgs) (*card.CardData, error) {

	cardNumberEnc, err := enc.EncryptByte([]byte(args.CardNumber))
	if err != nil {
		return nil, err
	}

	cardNameEnc, err := enc.EncryptByte([]byte(args.CardName))
	if err != nil {
		return nil, err
	}

	cardCVCEnc, err := enc.EncryptByte([]byte(args.CVC))
	if err != nil {
		return nil, err
	}

	cardDateEnc, err := enc.EncryptByte([]byte(args.CardDate))
	if err != nil {
		return nil, err
	}

	cardCommentEnc, err := enc.EncryptByte([]byte(args.Comment))
	if err != nil {
		return nil, err
	}

	return card.New(
		converter.StringToUUID(args.UserID),
		cardNumberEnc,
		cardNameEnc,
		cardCVCEnc,
		cardDateEnc,
		cardCommentEnc,
	)
}

func ToEncUpdateCard(enc *encryption.Enc, args *UpdateCardArgs) (*card.CardData, error) {

	cardNumberEnc, err := enc.EncryptByte([]byte(args.CardNumber))
	if err != nil {
		return nil, err
	}

	cardNameEnc, err := enc.EncryptByte([]byte(args.CardName))
	if err != nil {
		return nil, err
	}

	cardCVCEnc, err := enc.EncryptByte([]byte(args.CVC))
	if err != nil {
		return nil, err
	}

	cardDateEnc, err := enc.EncryptByte([]byte(args.CardDate))
	if err != nil {
		return nil, err
	}

	cardCommentEnc, err := enc.EncryptByte([]byte(args.Comment))
	if err != nil {
		return nil, err
	}

	date := time.Now().UTC()

	return card.NewWithID(
		converter.StringToUUID(args.ID),
		converter.StringToUUID(args.UserID),
		cardNumberEnc,
		cardNameEnc,
		cardCVCEnc,
		cardDateEnc,
		cardCommentEnc,
		date,
		date,
	)
}

func ToDecCard(enc *encryption.Enc, val *card.CardData) (*card.CardData, error) {

	cardNameDec, err := enc.DecryptByte(val.CardName())
	if err != nil {
		return nil, err
	}

	cardNumberDec, err := enc.DecryptByte(val.CardNumber())
	if err != nil {
		return nil, err
	}

	cardCVCDec, err := enc.DecryptByte(val.CVC())
	if err != nil {
		return nil, err
	}

	cardDateDec, err := enc.DecryptByte(val.CardDate())
	if err != nil {
		return nil, err
	}

	cardCommentDec, err := enc.DecryptByte(val.Comment())
	if err != nil {
		return nil, err
	}

	date := time.Now().UTC()

	return card.NewWithID(
		val.ID(),
		val.UserID(),
		cardNumberDec,
		cardNameDec,
		cardCVCDec,
		cardDateDec,
		cardCommentDec,
		date,
		date,
	)
}
