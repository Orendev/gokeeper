package account

import (
	"github.com/Orendev/gokeeper/pkg/type/url"
	"time"

	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Description of types

var (
	ErrLoginRequired    = errors.New("login is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrUserIDRequired   = errors.New("userID is required")
)

type Account struct {
	id        uuid.UUID
	userID    uuid.UUID
	title     title.Title
	login     login.Login
	password  password.Password
	url       url.URL
	comment   comment.Comment
	createdAt time.Time
	updatedAt time.Time
	isDeleted bool
}

// NewWithID - constructor a new instance of Account assets data with an ID.
func NewWithID(
	id uuid.UUID,
	userID uuid.UUID,
	title title.Title,
	login login.Login,
	password password.Password,
	url url.URL,
	comment comment.Comment,
	createdAt time.Time,
	updatedAt time.Time,
) (*Account, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	if login.IsEmpty() {
		return nil, ErrLoginRequired
	}

	if password.IsEmpty() {
		return nil, ErrPasswordRequired
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
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
	login login.Login,
	password password.Password,
	url url.URL,
	comment comment.Comment,
) (*Account, error) {

	if login.IsEmpty() {
		return nil, ErrLoginRequired
	}

	if password.IsEmpty() {
		return nil, ErrPasswordRequired
	}

	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}

	var timeNow = time.Now().UTC()

	return &Account{
		id:        uuid.New(),
		userID:    userID,
		title:     title,
		login:     login,
		password:  password,
		url:       url,
		comment:   comment,
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
func (a *Account) Login() login.Login {
	return a.login
}

// Password getter for the field
func (a *Account) Password() password.Password {
	return a.password
}

// URL getter for the field
func (a *Account) URL() url.URL {
	return a.url
}

// Comment getter for the field
func (a *Account) Comment() comment.Comment {
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

// Equal compare two accounts
func (a *Account) Equal(account Account) bool {
	return a.id == account.id
}
