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

	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
}

var CreateColumnUser = []string{
	"id",
	"created_at",
	"updated_at",
	"password",
	"email",
	"name",
	"surname",
	"patronymic",
}
