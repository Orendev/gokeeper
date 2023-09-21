package sqlite

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/pkg/errors"
)

func (r *Repository) CreateUser(ctx context.Context, user user.User) (*user.User, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	response, err := r.createUserTx(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r Repository) createUserTx(ctx context.Context, tx *sql.Tx, user user.User) (*user.User, error) {

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO user (id, name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			return
		}
	}()

	_, err = stmt.ExecContext(ctx, user.ID(), user.Name(), user.Email().String(), user.Password(), user.Role().String(), user.CreatedAt(), user.UpdatedAt())

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FindUser(ctx context.Context, email email.Email) (*user.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.findUserTx(ctx, tx, email)
}

func (r Repository) findUserTx(ctx context.Context, tx *sql.Tx, email email.Email) (*user.User, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"updated_at",
		"email",
		"role",
		"name",
	).From("user")

	builder = builder.Where(squirrel.Eq{"email": email.String()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoUser []*dao.User
	if err = sqlscan.ScanAll(&daoUser, rows); err != nil {
		return nil, err
	}

	if len(daoUser) == 0 {
		return nil, errors.New("user not found")
	}

	return r.toDomainUser(daoUser[0])
}

func (r Repository) LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	panic("implement me")
}
