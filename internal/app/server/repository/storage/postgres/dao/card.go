package dao

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID         uuid.UUID `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	UserId     uuid.UUID `db:"user_id"`
	CardNumber []byte    `db:"card_number"`
	CardName   []byte    `db:"card_name"`
	CardDate   []byte    `db:"card_date"`
	CVV        []byte    `db:"cvv"`
	Comment    []byte    `db:"comment"`
}

var CreateColumnCard = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"card_number",
	"card_name",
	"card_date",
	"cvv",
	"comment",
}
