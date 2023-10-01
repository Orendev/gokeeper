package grpc

import (
	"time"

	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/dto"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"
	"github.com/google/uuid"
)

func toAccountId(value string) (uuid.UUID, error) {

	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func toDomainAccount(dto dto.Account) (*account.Account, error) {

	titleObj, err := title.New(dto.Title)
	if err != nil {
		return nil, err
	}

	loginObj, err := login.New(dto.Login)
	if err != nil {
		return nil, err
	}

	passwordObj, err := password.New(dto.Password)
	if err != nil {
		return nil, err
	}

	urlObj, err := url.New(dto.URL)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, dto.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse(time.RFC3339, dto.UpdatedAt)
	if err != nil {
		return nil, err
	}

	commentObj, err := comment.New(dto.Comment)
	if err != nil {
		return nil, err
	}

	result, err := account.NewWithID(
		id,
		*titleObj,
		*loginObj,
		*passwordObj,
		*urlObj,
		*commentObj,
		dto.IsDeleted,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func toDomainAccounts(dto dto.ListAccount) ([]*account.Account, error) {

	result := make([]*account.Account, len(dto.Data))

	for i, val := range dto.Data {
		a, err := toDomainAccount(val)
		if err != nil {
			return nil, err
		}

		result[i] = a
	}

	return result, nil
}
