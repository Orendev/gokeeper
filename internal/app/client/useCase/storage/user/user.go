package user

import (
	"context"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/google/uuid"

	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
)

// Add creating a user
func (uc *UseCase) Add(ctx context.Context, user user.User) (*user.User, error) {
	return uc.adapterStorage.AddUser(ctx, user)
}

// UpdateToken update token a user
func (uc *UseCase) UpdateToken(ctx context.Context, id uuid.UUID, token token.Token) (*user.User, error) {
	return uc.adapterStorage.UpdateToken(ctx, id, token)
}

// Get let's get the user
func (uc *UseCase) Get(ctx context.Context) (*user.User, error) {
	return uc.adapterStorage.GetUser(ctx)
}
