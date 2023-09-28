package sqlite

import (
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/comment"
	"github.com/Orendev/gokeeper/pkg/type/login"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/Orendev/gokeeper/pkg/type/url"

	"github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite/dao"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) toCopyFromSourceAccounts(accounts ...*account.Account) pgx.CopyFromSource {
	rows := make([][]interface{}, len(accounts))

	for i, val := range accounts {
		rows[i] = []interface{}{
			val.ID(),
			val.Password().String(),
			val.Login().String(),
			val.Title().String(),
			val.URL().String(),
			val.Comment().String(),
			val.IsDeleted(),
			val.CreatedAt(),
			val.UpdatedAt(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r *Repository) toDomainAccount(dao *dao.Account) (*account.Account, error) {

	titleObj, err := title.New(dao.Title.String)
	if err != nil {
		return nil, err
	}

	loginObj, err := login.New(dao.Login)
	if err != nil {
		return nil, err
	}

	passwordObj, err := password.New(dao.Password)
	if err != nil {
		return nil, err
	}

	urlObj, err := url.New(dao.URL.String)
	if err != nil {
		return nil, err
	}

	commentObj, err := comment.New(dao.Comment.String)
	if err != nil {
		return nil, err
	}

	result, err := account.NewWithID(
		dao.ID,
		*titleObj,
		*loginObj,
		*passwordObj,
		*urlObj,
		*commentObj,
		dao.IsDeleted,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) toDomainAccounts(dao []*dao.Account) ([]*account.Account, error) {
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
