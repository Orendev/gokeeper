package account

import (
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrLoginRequired    = errors.New("login is required")
	ErrPasswordRequired = errors.New("password is required")
)

type Account struct {
	id        uuid.UUID
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
	title title.Title,
	login login.Login,
	password password.Password,
	url url.URL,
	comment comment.Comment,
	isDeleted bool,
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

	return &Account{
		id:        id,
		title:     title,
		login:     login,
		password:  password,
		url:       url,
		comment:   comment,
		isDeleted: isDeleted,
		createdAt: createdAt.UTC(),
		updatedAt: updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of Account.
func New(
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

	var timeNow = time.Now().UTC()

	return &Account{
		id:        uuid.New(),
		title:     title,
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
