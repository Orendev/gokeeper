package dao

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	UserId uuid.UUID `db:"user_id"`

	Login    string `db:"login"`
	Password string `db:"password"`

	Title      string `db:"title"`
	Comment    string `db:"comment"`
	WebAddress string `db:"web_address"`

	Version uint64 `db:"version"`
}

var CreateColumnAccount = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"login",
	"password",
	"title",
	"comment",
	"web_address",
	"version",
}
