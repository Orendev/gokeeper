package user

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
)

func (uc *UseCase) Create(users ...*user.User) ([]*user.User, error) {
	return uc.adapterStorage.CreateUser(users...)
}

func (uc *UseCase) Find(email email.Email) (*user.User, error) {
	return uc.adapterStorage.FindUser(email)
}
