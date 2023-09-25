package user

import (
	"github.com/Orendev/gokeeper/internal/app/client/useCase/adapters/storage"
)

type UseCase struct {
	adapterStorage storage.User
	options        Options
}
type Options struct {
}

func New(storage storage.User, options Options) *UseCase {
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
