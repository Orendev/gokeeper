package card

import (
	"github.com/Orendev/gokeeper/internal/app/server/useCase/adapters/storage"
)

type UseCase struct {
	adapterStorage storage.Card
	options        Options
}

type Options struct{}

func New(storage storage.Card, options Options) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}