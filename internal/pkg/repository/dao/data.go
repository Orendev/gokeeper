package dao

import (
	"time"

	"github.com/google/uuid"
)

var TableNameText = "texts"
var TableNameBinary = "binaries"

// Data description of fields in the database
type Data struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserId    uuid.UUID `db:"user_id"`
	Title     string    `db:"title"`
	Data      []byte    `db:"data"`
	Comment   []byte    `db:"comment"`
	IsDeleted bool      `db:"is_deleted"`
}

// ColumnData Names of fields in the database
var ColumnData = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"title",
	"data",
	"comment",
	"is_deleted",
}
