package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/jackc/pgx/v5"
)

var mappingSortCard = map[columnCode.ColumnCode]string{
	"id": "id",
}

func (r Repository) toCopyFromSourceCards(cards ...*card.CardData) pgx.CopyFromSource {
	rows := make([][]interface{}, len(cards))

	for i, val := range cards {

		rows[i] = []interface{}{
			val.ID().String(),
			val.CreatedAt().UTC(),
			val.UpdatedAt().UTC(),
			val.UserID().String(),
			val.CardNumber(),
			val.CardName(),
			val.CardDate(),
			val.CVC(),
			val.Comment(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainCard(dao *dao.Card) (*card.CardData, error) {

	return card.NewWithID(
		dao.ID,
		dao.UserId,
		dao.CardNumber,
		dao.CardName,
		dao.CVV,
		dao.CardDate,
		dao.Comment,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
}

func (r Repository) toDomainCards(dao []*dao.Card) ([]*card.CardData, error) {
	var result = make([]*card.CardData, len(dao))
	for i, v := range dao {
		c, err := r.toDomainCard(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
