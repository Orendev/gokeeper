package text

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(texts ...*text.TextData) ([]*text.TextData, error) {
	return uc.adapterStorage.CreateText(texts...)
}

func (uc *UseCase) Update(textUpdate text.TextData) (*text.TextData, error) {
	return uc.adapterStorage.UpdateText(textUpdate.ID(), func(oldText *text.TextData) (*text.TextData, error) {
		return text.NewWithID(
			oldText.ID(),
			oldText.UserID(),
			textUpdate.Title(),
			textUpdate.Body(),
			textUpdate.Comment(),
			textUpdate.Version(),
			oldText.CreatedAt(),
			time.Now().UTC(),
		)
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteText(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
	return uc.adapterStorage.ListText(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *text.TextData, err error) {
	return uc.adapterStorage.ReadTextByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountText()
}
