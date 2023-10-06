package dao

import (
	"time"

	"github.com/google/uuid"
)

var TableNameCard = "cards"

type Card struct {
	ID         uuid.UUID `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	UserID     uuid.UUID `db:"user_id"`
	CardNumber []byte    `db:"card_number"`
	CardName   []byte    `db:"card_name"`
	CardDate   []byte    `db:"card_date"`
	CVV        []byte    `db:"cvv"`
	Comment    []byte    `db:"comment"`
	IsDelete   bool      `db:"is_deleted"`
}

var ColumnCard = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"card_number",
	"card_name",
	"card_date",
	"cvc",
	"comment",
	"is_deleted",
}
