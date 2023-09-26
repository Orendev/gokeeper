package user

import (
	"context"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"time"
)

// Add creating a user
func (uc *UseCase) Add(ctx context.Context, user user.User) (*user.User, error) {
	return uc.adapterStorage.AddUser(ctx, user)
}

// UpdateToken update token a user
func (uc *UseCase) UpdateToken(ctx context.Context, update user.User) (*user.User, error) {
	return uc.adapterStorage.UpdateToken(ctx, update.ID(), func(old *user.User) (*user.User, error) {
		return user.NewWithID(old.ID(), update.Password(), update.Email(), update.Role(), update.Name(), update.Token(), old.CreatedAt(), time.Now().UTC())
	})
}

// Get let's get the user
func (uc *UseCase) Get(ctx context.Context) (*user.User, error) {
	return uc.adapterStorage.GetUser(ctx)
}
