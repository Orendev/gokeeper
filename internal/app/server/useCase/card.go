package useCase

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Card interface {
	Create(cards ...*card.CardData) ([]*card.CardData, error)
	Update(card card.CardData) (*card.CardData, error)
	Delete(ID uuid.UUID) error

	CardReader
}

type CardReader interface {
	List(parameter queryParameter.QueryParameter) ([]*card.CardData, error)
	ReadByID(ID uuid.UUID) (response *card.CardData, err error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
