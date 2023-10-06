package card

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrNumberRequired = errors.New("number is required")
	ErrCVCRequired    = errors.New("CVC is required")
	ErrUserIDRequired = errors.New("userID is required")
)

type ListCardViewModel struct {
	Data   []*CardData `json:"data"`
	Total  uint64      `json:"total"`
	Limit  uint64      `json:"limit"`
	Offset uint64      `json:"offset"`
}

type CardData struct {
	id         uuid.UUID
	userID     uuid.UUID
	cardNumber []byte
	cardName   []byte
	cvc        []byte
	cardDate   []byte
	comment    []byte
	createdAt  time.Time
	updatedAt  time.Time
	isDeleted  bool
}

// NewWithID - constructor a new instance of Account assets data with an ID.
func NewWithID(
	id uuid.UUID,
	userID uuid.UUID,
	cardNumber []byte,
	cardName []byte,
	cvc []byte,
	cardDate []byte,
	comment []byte,
	createdAt time.Time,
	updatedAt time.Time,
) (*CardData, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	if len(cardNumber) == 0 {
		return nil, ErrNumberRequired
	}

	if len(cvc) == 0 {
		return nil, ErrCVCRequired
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	return &CardData{
		id:         id,
		userID:     userID,
		cardNumber: cardNumber,
		cardName:   cardName,
		cvc:        cvc,
		cardDate:   cardDate,
		comment:    comment,
		createdAt:  createdAt.UTC(),
		updatedAt:  updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of Account.
func New(
	userID uuid.UUID,
	cardNumber []byte,
	cardName []byte,
	cvc []byte,
	cardDate []byte,
	comment []byte,
) (*CardData, error) {

	if len(cardNumber) == 0 {
		return nil, ErrNumberRequired
	}

	if len(cvc) == 0 {
		return nil, ErrCVCRequired
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	var timeNow = time.Now().UTC()

	return &CardData{
		id:         uuid.New(),
		userID:     userID,
		cardNumber: cardNumber,
		cardName:   cardName,
		cvc:        cvc,
		cardDate:   cardDate,
		comment:    comment,

		createdAt: timeNow,
		updatedAt: timeNow,
	}, nil
}

// ID getter for the field
func (d CardData) ID() uuid.UUID {
	return d.id
}

// UserID getter for the field
func (d CardData) UserID() uuid.UUID {
	return d.userID
}

// CardNumber getter for the field
func (d CardData) CardNumber() []byte {
	return d.cardNumber
}

// CardName getter for the field
func (d CardData) CardName() []byte {
	return d.cardName
}

// CVC getter for the field
func (d CardData) CVC() []byte {
	return d.cvc
}

// CardDate getter for the field
func (d CardData) CardDate() []byte {
	return d.cardDate
}

// Comment getter for the field
func (d CardData) Comment() []byte {
	return d.comment
}

// CreatedAt getter for the field
func (d CardData) CreatedAt() time.Time {
	return d.createdAt
}

// UpdatedAt getter for the field
func (d CardData) UpdatedAt() time.Time {
	return d.updatedAt
}

// IsDeleted getter for the field
func (d CardData) IsDeleted() bool {
	return d.isDeleted
}
