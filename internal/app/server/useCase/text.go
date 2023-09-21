package useCase

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interaction between delivery and use case

type Text interface {
	Create(texts ...*text.TextData) ([]*text.TextData, error)
	Update(text text.TextData) (*text.TextData, error)
	Delete(ID uuid.UUID) error

	TextReader
}

type TextReader interface {
	List(parameter queryParameter.QueryParameter) ([]*text.TextData, error)
	ReadByID(ID uuid.UUID) (response *text.TextData, err error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
