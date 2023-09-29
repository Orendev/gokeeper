package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"
	"github.com/jackc/pgx/v5"
)

func (r Repository) toCopyFromSourceAccounts(accounts ...*account.Account) pgx.CopyFromSource {
	rows := make([][]interface{}, len(accounts))

	for i, val := range accounts {

		rows[i] = []interface{}{
			val.ID().String(),
			val.CreatedAt().UTC(),
			val.UpdatedAt().UTC(),
			val.UserID().String(),
			val.Login().String(),
			val.Password().String(),
			val.Title().String(),
			val.Comment().String(),
			val.URL().String(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainAccount(dao *dao.Account) (*account.Account, error) {

	urlObject, err := url.New(dao.URL)
	if err != nil {
		return nil, err
	}

	titleObject, err := title.New(dao.Title)
	if err != nil {
		return nil, err
	}

	passwordObject, err := password.New(dao.Password)
	if err != nil {
		return nil, err
	}

	loginObject, err := login.New(dao.Login)
	if err != nil {
		return nil, err
	}

	commentObject, err := comment.New(dao.Comment)
	if err != nil {
		return nil, err
	}

	return account.NewWithID(
		dao.ID,
		dao.UserId,
		*titleObject,
		*loginObject,
		*passwordObject,
		*urlObject,
		*commentObject,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
}

func (r Repository) toDomainAccounts(dao []*dao.Account) ([]*account.Account, error) {
	var result = make([]*account.Account, len(dao))
	for i, v := range dao {
		c, err := r.toDomainAccount(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
