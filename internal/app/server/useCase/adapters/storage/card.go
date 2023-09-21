package storage

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Card interface {
	CreateCard(cards ...*card.CardData) ([]*card.CardData, error)
	UpdateCard(ID uuid.UUID, updateFn func(c *card.CardData) (*card.CardData, error)) (*card.CardData, error)
	DeleteCard(ID uuid.UUID) error

	CardReader
}

type CardReader interface {
	ListCard(parameter queryParameter.QueryParameter) ([]*card.CardData, error)
	ReadCardByID(ID uuid.UUID) (response *card.CardData, err error)
	CountCard( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
