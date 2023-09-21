package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

func (r *Repository) CreateCard(cards ...*card.CardData) ([]*card.CardData, error) {
	panic("implement me")
}

func (r *Repository) UpdateCard(ID uuid.UUID, updateFn func(a *card.CardData) (*card.CardData, error)) (*card.CardData, error) {
	panic("implement me")
}

func (r *Repository) DeleteCard(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListCard(parameter queryParameter.QueryParameter) ([]*card.CardData, error) {
	panic("implement me")
}

func (r *Repository) ReadCardByID(ID uuid.UUID) (response *card.CardData, err error) {
	panic("implement me")
}

func (r *Repository) CountCard() (uint64, error) {
	panic("implement me")
}
