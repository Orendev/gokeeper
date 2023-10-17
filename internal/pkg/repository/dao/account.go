package dao

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
)

var TableNameAccount = "accounts"

var SortAccount = map[columnCode.ColumnCode]string{
	"id": "id",
}

type Account struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserID    uuid.UUID `db:"user_id"`
	Password  []byte    `db:"password"`
	Login     []byte    `db:"login"`
	Title     string    `db:"title"`
	Comment   []byte    `db:"comment"`

	URL       []byte `db:"url"`
	IsDeleted bool   `db:"is_deleted"`
}

var ColumnAccount = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"password",
	"login",
	"title",
	"url",
	"comment",
	"is_deleted",
}

func ToDomainAccount(dao *Account) (*account.Account, error) {

	titleObj, err := title.New(dao.Title)
	if err != nil {
		return nil, err
	}

	return account.NewWithID(
		dao.ID,
		dao.UserID,
		*titleObj,
		dao.Login,
		dao.Password,
		dao.URL,
		dao.Comment,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
}

func ToDomainAccounts(dao []*Account) ([]*account.Account, error) {
	var result = make([]*account.Account, len(dao))
	for i, v := range dao {
		c, err := ToDomainAccount(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}

	return result, nil
}
