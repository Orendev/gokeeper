package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user/name"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user/patronymic"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user/surname"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/jackc/pgx/v4"
)

func (r Repository) toCopyFromSourceUsers(users ...*user.User) pgx.CopyFromSource {
	rows := make([][]interface{}, len(users))

	for i, val := range users {
		rows[i] = []interface{}{
			val.ID(),
			val.Password().String(),
			val.Email().String(),
			val.Name().String(),
			val.Surname().String(),
			val.Patronymic().String(),
			val.CreatedAt(),
			val.UpdatedAt(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainUser(dao *dao.User) (*user.User, error) {

	emailObject, err := email.New(dao.Email)
	if err != nil {
		return nil, err
	}

	passwordObject, err := password.New(dao.Password)
	if err != nil {
		return nil, err
	}

	nameObject, err := name.New(dao.Name)
	if err != nil {
		return nil, err
	}

	surnametObject, err := surname.New(dao.Surname)
	if err != nil {
		return nil, err
	}

	patronymicObject, err := patronymic.New(dao.Patronymic)
	if err != nil {
		return nil, err
	}

	result, err := user.NewWithID(
		dao.ID,
		*passwordObject,
		*emailObject,
		*nameObject,
		*surnametObject,
		*patronymicObject,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r Repository) toDomainUsers(dao []*dao.User) ([]*user.User, error) {
	var result = make([]*user.User, len(dao))
	for i, v := range dao {
		c, err := r.toDomainUser(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
