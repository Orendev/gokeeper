package text

import (
	"time"

	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrUserIDRequired = errors.New("userID is required")

type TextData struct {
	id        uuid.UUID
	userID    uuid.UUID
	title     title.Title
	data      []byte
	comment   []byte
	createdAt time.Time
	updatedAt time.Time
	isDeleted bool
}

// NewWithID - constructor a new instance of TextData assets data with an ID.
func NewWithID(
	id uuid.UUID,
	userID uuid.UUID,
	title title.Title,
	data []byte,
	comment []byte,
	createdAt time.Time,
	updatedAt time.Time,
) (*TextData, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	return &TextData{
		id:        id,
		userID:    userID,
		title:     title,
		data:      data,
		comment:   comment,
		createdAt: createdAt.UTC(),
		updatedAt: updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of Account.
func New(
	userID uuid.UUID,
	title title.Title,
	data []byte,
	comment []byte,
) (*TextData, error) {

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	var timeNow = time.Now().UTC()

	return &TextData{
		id:        uuid.New(),
		userID:    userID,
		title:     title,
		data:      data,
		comment:   comment,
		createdAt: timeNow,
		updatedAt: timeNow,
	}, nil
}

// ID getter for the field
func (d *TextData) ID() uuid.UUID {
	return d.id
}

// UserID getter for the field
func (d *TextData) UserID() uuid.UUID {
	return d.userID
}

// Title getter for the field
func (d *TextData) Title() title.Title {
	return d.title
}

// Data getter for the field
func (d *TextData) Data() []byte {
	return d.data
}

// Comment getter for the field
func (d *TextData) Comment() []byte {
	return d.comment
}

// CreatedAt getter for the field
func (d *TextData) CreatedAt() time.Time {
	return d.createdAt
}

// UpdatedAt getter for the field
func (d *TextData) UpdatedAt() time.Time {
	return d.updatedAt
}

// IsDeleted getter for the field
func (d *TextData) IsDeleted() bool {
	return d.isDeleted
}
