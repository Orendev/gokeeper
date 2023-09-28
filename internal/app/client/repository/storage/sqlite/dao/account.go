package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Password string         `db:"password"`
	Login    string         `db:"login"`
	Title    sql.NullString `db:"title"`
	Comment  sql.NullString `db:"comment"`

	URL       sql.NullString `db:"url"`
	IsDeleted bool           `db:"is_deleted"`
}

var ColumnAccount = []string{
	"id",
	"created_at",
	"updated_at",
	"password",
	"login",
	"title",
	"url",
	"comment",
	"is_deleted",
}
