package binary

import (
	"time"

	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrUserIDRequired = errors.New("userID is required")

type ListBinaryViewModel struct {
	Data   []*BinaryData `json:"data"`
	Total  uint64        `json:"total"`
	Limit  uint64        `json:"limit"`
	Offset uint64        `json:"offset"`
}

type BinaryData struct {
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
) (*BinaryData, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	return &BinaryData{
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
) (*BinaryData, error) {

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	var timeNow = time.Now().UTC()

	return &BinaryData{
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
func (d *BinaryData) ID() uuid.UUID {
	return d.id
}

// UserID getter for the field
func (d *BinaryData) UserID() uuid.UUID {
	return d.userID
}

// Title getter for the field
func (d *BinaryData) Title() title.Title {
	return d.title
}

// Data getter for the field
func (d *BinaryData) Data() []byte {
	return d.data
}

// Comment getter for the field
func (d *BinaryData) Comment() []byte {
	return d.comment
}

// CreatedAt getter for the field
func (d *BinaryData) CreatedAt() time.Time {
	return d.createdAt
}

// UpdatedAt getter for the field
func (d *BinaryData) UpdatedAt() time.Time {
	return d.updatedAt
}

// IsDeleted getter for the field
func (d *BinaryData) IsDeleted() bool {
	return d.isDeleted
}
