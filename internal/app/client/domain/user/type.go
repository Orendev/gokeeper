package user

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user/name"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user/role"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/hashedPassword"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User is representing the User data struct
type User struct {
	id        uuid.UUID
	password  password.Password
	email     email.Email
	name      name.Name
	role      role.Role
	token     token.Token
	createdAt time.Time
	updatedAt time.Time
}

// NewWithID - constructor a new instance of User assets data with an ID.
func NewWithID(
	id uuid.UUID,
	password password.Password,
	email email.Email,
	role role.Role,
	name name.Name,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {

	if id == uuid.Nil {
		id = uuid.New()
	}

	return &User{
		id:        id,
		password:  password,
		email:     email,
		name:      name,
		role:      role,
		createdAt: createdAt.UTC(),
		updatedAt: updatedAt.UTC(),
	}, nil
}

// New - constructor a new instance of User.
func New(
	password password.Password,
	email email.Email,
	name name.Name,
) (*User, error) {

	var timeNow = time.Now().UTC()

	return &User{
		id:        uuid.New(),
		password:  password,
		email:     email,
		name:      name,
		role:      role.User,
		createdAt: timeNow,
		updatedAt: timeNow,
	}, nil
}

// ID getter for the field
func (u User) ID() uuid.UUID {
	return u.id
}

// Password getter for the field
func (u User) Password() password.Password {
	return u.password
}

// Email getter for the field
func (u User) Email() email.Email {
	return u.email
}

// Name getter for the field
func (u User) Name() name.Name {
	return u.name
}

// Role getter for the field
func (u User) Role() role.Role {
	return u.role
}

// CreatedAt getter for the field
func (u User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt getter for the field
func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

// Equal compare two accounts
func (u User) Equal(user User) bool {
	return u.id == user.id
}

func (u User) IsCorrectPassword(hashedPassword hashedPassword.HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword.Byte(), u.password.Byte())
	return err == nil
}
