package card

import (
	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
)

var SortData = map[columnCode.ColumnCode]string{
	"id": "id",
}

func ToDomainCard(dao *dao.Card) (*card.CardData, error) {

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

func ToDomainCards(dao []*dao.Card) ([]*card.CardData, error) {
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
