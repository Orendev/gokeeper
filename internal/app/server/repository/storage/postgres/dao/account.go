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

	Title   string `db:"title"`
	Comment string `db:"comment"`
	URL     string `db:"url"`
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
	"url",
}
