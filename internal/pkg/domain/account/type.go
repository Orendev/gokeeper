package account

import (
	"time"

	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
)

type ListAccountViewModel struct {
	Data   []*Account `json:"data"`
	Total  uint64     `json:"total"`
	Limit  uint64     `json:"limit"`
	Offset uint64     `json:"offset"`
}

type Account struct {
	id        uuid.UUID
	userID    uuid.UUID
	title     title.Title
	login     []byte
	password  []byte
	url       []byte
	comment   []byte
	createdAt time.Time
	updatedAt time.Time
	isDeleted bool
}

// NewWithID - constructor a new instance of Account assets data with an ID.
func NewWithID(
	id uuid.UUID,
	userID uuid.UUID,
	title title.Title,
	login []byte,
	password []byte,
	url []byte,
	comment []byte,
	createdAt time.Time,
	updatedAt time.Time,
) (*Account, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	return &Account{
		id:        id,
		userID:    userID,
		title:     title,
		login:     login,
		password:  password,
		url:       url,
		comment:   comment,
		createdAt: createdAt.UTC(),
		updatedAt: updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of Account.
func New(
	userID uuid.UUID,
	title title.Title,
	login []byte,
	password []byte,
	url []byte,
	comment []byte,
) (*Account, error) {

	var timeNow = time.Now().UTC()

	return &Account{
		id:        uuid.New(),
		title:     title,
		userID:    userID,
		login:     login,
		password:  password,
		url:       url,
		comment:   comment,
		isDeleted: false,
		createdAt: timeNow,
		updatedAt: timeNow,
	}, nil
}

// ID getter for the field
func (a *Account) ID() uuid.UUID {
	return a.id
}

// UserID getter for the field
func (a *Account) UserID() uuid.UUID {
	return a.userID
}

// Title getter for the field
func (a *Account) Title() title.Title {
	return a.title
}

// Login getter for the field
func (a *Account) Login() []byte {
	return a.login
}

// Password getter for the field
func (a *Account) Password() []byte {
	return a.password
}

// URL getter for the field
func (a *Account) URL() []byte {
	return a.url
}

// Comment getter for the field
func (a *Account) Comment() []byte {
	return a.comment
}

// CreatedAt getter for the field
func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}

// UpdatedAt getter for the field
func (a *Account) UpdatedAt() time.Time {
	return a.updatedAt
}

// IsDeleted getter for the field
func (a *Account) IsDeleted() bool {
	return a.isDeleted
}
