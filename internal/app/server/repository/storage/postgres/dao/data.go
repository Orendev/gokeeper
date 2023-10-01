package dao

import (
	"time"

	"github.com/google/uuid"
)

type Data struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserId    uuid.UUID `db:"user_id"`
	Title     string    `db:"title"`
	Data      []byte    `db:"data"`
	Comment   []byte    `db:"comment"`
}

var CreateColumnData = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"title",
	"data",
	"comment",
}
