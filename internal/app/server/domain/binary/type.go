package binary

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/binary/body"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/version"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrUserIDRequired = errors.New("userID is required")

type BinaryData struct {
	id        uuid.UUID
	userID    uuid.UUID
	title     title.Title
	body      body.Body
	comment   comment.Comment
	version   version.Version
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

// NewWithID - constructor a new instance of TextData assets data with an ID.
func NewWithID(
	id uuid.UUID,
	userID uuid.UUID,
	title title.Title,
	body body.Body,
	comment comment.Comment,
	version version.Version,
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
		body:      body,
		comment:   comment,
		version:   version,
		createdAt: createdAt.UTC(),
		updatedAt: updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of Account.
func New(
	userID uuid.UUID,
	title title.Title,
	body body.Body,
	comment comment.Comment,
	version version.Version,
) (*BinaryData, error) {

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	var timeNow = time.Now().UTC()

	return &BinaryData{
		id:        uuid.New(),
		userID:    userID,
		title:     title,
		body:      body,
		comment:   comment,
		version:   version,
		createdAt: timeNow,
		updatedAt: timeNow,
	}, nil
}

// ID getter for the field
func (d BinaryData) ID() uuid.UUID {
	return d.id
}

// UserID getter for the field
func (d BinaryData) UserID() uuid.UUID {
	return d.userID
}

// Title getter for the field
func (d BinaryData) Title() title.Title {
	return d.title
}

// Body getter for the field
func (d BinaryData) Body() body.Body {
	return d.body
}

// Comment getter for the field
func (d BinaryData) Comment() comment.Comment {
	return d.comment
}

// Version getter for the field
func (d BinaryData) Version() version.Version {
	return d.version
}

// CreatedAt getter for the field
func (d BinaryData) CreatedAt() time.Time {
	return d.createdAt
}

// UpdatedAt getter for the field
func (d BinaryData) UpdatedAt() time.Time {
	return d.updatedAt
}

// DeletedAt getter for the field
func (d BinaryData) DeletedAt() time.Time {
	return d.deletedAt
}

// Equal compare two accounts
func (d BinaryData) Equal(binaryData BinaryData) bool {
	return d.id == binaryData.id
}
