package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/internal/app/server/domain/account/webAddress"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/version"
	"github.com/jackc/pgx/v4"
)

func (r Repository) toCopyFromSourceAccounts(accounts ...*account.Account) pgx.CopyFromSource {
	rows := make([][]interface{}, len(accounts))

	for i, val := range accounts {
		rows[i] = []interface{}{
			val.ID(),
			val.UserID(),
			val.Title().String(),
			val.Login().String(),
			val.Password().String(),
			val.WebAddress().String(),
			val.Comment(),
			val.Version().String(),
			val.CreatedAt(),
			val.UpdatedAt(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainAccount(dao *dao.Account) (*account.Account, error) {

	webAddressObject, err := webAddress.New(dao.WebAddress)
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

	versionObject, err := version.New(dao.Version)
	if err != nil {
		return nil, err
	}

	result, err := account.NewWithID(
		dao.ID,
		dao.UserId,
		*titleObject,
		*loginObject,
		*passwordObject,
		*webAddressObject,
		*commentObject,
		*versionObject,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
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
