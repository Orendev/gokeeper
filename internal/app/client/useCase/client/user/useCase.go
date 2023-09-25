package user

import (
	"github.com/Orendev/gokeeper/internal/app/client/useCase/adapters/client"
)

type UseCase struct {
	adapterClient client.User
	options       Options
}

type Options struct {
}

func New(client client.User, options Options) *UseCase {
	var uc = &UseCase{
		adapterClient: client,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
