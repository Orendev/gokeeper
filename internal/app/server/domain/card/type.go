package card

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/card/cvc"
	"github.com/Orendev/gokeeper/internal/app/server/domain/card/name"
	"github.com/Orendev/gokeeper/internal/app/server/domain/card/number"
	"github.com/Orendev/gokeeper/internal/app/server/domain/card/surname"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/version"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrNumberRequired         = errors.New("number is required")
	ErrExpirationDateRequired = errors.New("expirationDate is required")
	ErrCVCRequired            = errors.New("CVC is required")
	ErrUserIDRequired         = errors.New("userID is required")
)

type CardData struct {
	id             uuid.UUID
	userID         uuid.UUID
	number         number.Number
	name           name.Name
	surname        surname.Surname
	cvc            cvc.CVC
	expirationDate time.Time
	comment        comment.Comment
	version        version.Version
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      time.Time
}

// NewWithID - constructor a new instance of Account assets data with an ID.
func NewWithID(
	id uuid.UUID,
	userID uuid.UUID,
	number number.Number,
	name name.Name,
	surname surname.Surname,
	cvc cvc.CVC,
	expirationDate time.Time,
	comment comment.Comment,
	version version.Version,
	createdAt time.Time,
	updatedAt time.Time,
) (*CardData, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	if number.IsEmpty() {
		return nil, ErrNumberRequired
	}

	if cvc.IsEmpty() {
		return nil, ErrCVCRequired
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	return &CardData{
		id:             id,
		userID:         userID,
		number:         number,
		name:           name,
		surname:        surname,
		cvc:            cvc,
		expirationDate: expirationDate,
		comment:        comment,
		version:        version,
		createdAt:      createdAt.UTC(),
		updatedAt:      updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of Account.
func New(
	userID uuid.UUID,
	number number.Number,
	name name.Name,
	surname surname.Surname,
	cvc cvc.CVC,
	expirationDate time.Time,
	comment comment.Comment,
	version version.Version,
) (*CardData, error) {

	if number.IsEmpty() {
		return nil, ErrNumberRequired
	}

	if cvc.IsEmpty() {
		return nil, ErrCVCRequired
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	var timeNow = time.Now().UTC()

	return &CardData{
		id:             uuid.New(),
		userID:         userID,
		number:         number,
		name:           name,
		surname:        surname,
		cvc:            cvc,
		expirationDate: expirationDate,
		comment:        comment,
		version:        version,
		createdAt:      timeNow,
		updatedAt:      timeNow,
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

// Number getter for the field
func (d CardData) Number() number.Number {
	return d.number
}

// Name getter for the field
func (d CardData) Name() name.Name {
	return d.name
}

// Surname getter for the field
func (d CardData) Surname() surname.Surname {
	return d.surname
}

// CVC getter for the field
func (d CardData) CVC() cvc.CVC {
	return d.cvc
}

// ExpirationDate getter for the field
func (d CardData) ExpirationDate() time.Time {
	return d.expirationDate
}

// Version getter for the field
func (d CardData) Version() version.Version {
	return d.version
}

// Comment getter for the field
func (d CardData) Comment() comment.Comment {
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

// DeletedAt getter for the field
func (d CardData) DeletedAt() time.Time {
	return d.deletedAt
}

// Equal compare two accounts
func (d CardData) Equal(cardType CardData) bool {
	return d.id == cardType.id
}
