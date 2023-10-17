package account

import (
	"github.com/Orendev/gokeeper/internal/pkg/useCase/adapters"
)

type UseCase struct {
	adapterStorage adapters.Account
	options        Options
}

type Options struct {
}

func New(client adapters.Account, options Options) *UseCase {
	var uc = &UseCase{
		adapterStorage: client,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
