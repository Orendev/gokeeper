package storage

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
)

// Interface for interacting with the use case repository.

type Text interface {
	CreateText(texts ...*text.TextData) ([]*text.TextData, error)
	UpdateText(ID uuid.UUID, updateFn func(t *text.TextData) (*text.TextData, error)) (*text.TextData, error)
	DeleteText(ID uuid.UUID) error

	TextReader
}

type TextReader interface {
	ListText(parameter queryParameter.QueryParameter) ([]*text.TextData, error)
	ReadTextByID(ID uuid.UUID) (response *text.TextData, err error)
	CountText( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
