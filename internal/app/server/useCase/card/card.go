package card

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/card"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(cards ...*card.CardData) ([]*card.CardData, error) {
	return uc.adapterStorage.CreateCard(cards...)
}

func (uc *UseCase) Update(cardUpdate card.CardData) (*card.CardData, error) {
	return uc.adapterStorage.UpdateCard(cardUpdate.ID(), func(oldCard *card.CardData) (*card.CardData, error) {
		return card.NewWithID(
			oldCard.ID(),
			oldCard.UserID(),
			cardUpdate.Number(),
			cardUpdate.Name(),
			cardUpdate.Surname(),
			cardUpdate.CVC(),
			cardUpdate.ExpirationDate(),
			cardUpdate.Comment(),
			cardUpdate.Version(),
			oldCard.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteCard(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*card.CardData, error) {
	return uc.adapterStorage.ListCard(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *card.CardData, err error) {
	return uc.adapterStorage.ReadCardByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountCard()
}
