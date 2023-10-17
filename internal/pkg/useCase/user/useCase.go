package user

import (
	"github.com/Orendev/gokeeper/internal/pkg/useCase/adapters"
)

type UseCase struct {
	adapterStorage adapters.User
	options        Options
}

type Options struct{}

func New(storage adapters.User, options Options) *UseCase {
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
