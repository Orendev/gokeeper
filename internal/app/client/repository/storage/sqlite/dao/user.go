package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Password string         `db:"password"`
	Email    string         `db:"email"`
	Role     string         `db:"role"`
	Token    sql.NullString `db:"token"`

	Name string `db:"name"`
}

var ColumnUser = []string{
	"id",
	"created_at",
	"updated_at",
	"password",
	"email",
	"role",
	"name",
	"token",
}
