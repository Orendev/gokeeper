package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

func (r *Repository) CreateText(texts ...*text.TextData) ([]*text.TextData, error) {
	panic("implement me")
}

func (r *Repository) UpdateText(ID uuid.UUID, updateFn func(t *text.TextData) (*text.TextData, error)) (*text.TextData, error) {
	panic("implement me")
}

func (r *Repository) DeleteText(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListText(parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
	panic("implement me")
}

func (r *Repository) ReadTextByID(ID uuid.UUID) (response *text.TextData, err error) {
	panic("implement me")
}

func (r *Repository) CountText() (uint64, error) {
	panic("implement me")
}
