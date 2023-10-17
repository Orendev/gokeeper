package dao

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/google/uuid"
)

var TableNameCard = "cards"

var SortCard = map[columnCode.ColumnCode]string{
	"id": "id",
}

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

func ToDomainCard(dao *Card) (*card.CardData, error) {

	return card.NewWithID(
		dao.ID,
		dao.UserID,
		dao.CardNumber,
		dao.CardName,
		dao.CVV,
		dao.CardDate,
		dao.Comment,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
}

func ToDomainCards(dao []*Card) ([]*card.CardData, error) {
	var result = make([]*card.CardData, len(dao))
	for i, v := range dao {
		c, err := ToDomainCard(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
