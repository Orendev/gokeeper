package dao

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Password string `db:"password"`
	Email    string `db:"email"`
	Token    string `db:"token"`
	Role     string `db:"token"`

	Name string `db:"name"`
}

var CreateColumnUser = []string{
	"id",
	"created_at",
	"updated_at",
	"password",
	"email",
	"name",
	"role",
	"token",
}
