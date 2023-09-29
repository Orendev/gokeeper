package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/server/domain/user"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *Repository) CreateUser(users ...*user.User) ([]*user.User, error) {
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createUserTx(ctx, tx, users...)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createUserTx(ctx context.Context, tx pgx.Tx, users ...*user.User) ([]*user.User, error) {
	if len(users) == 0 {
		return []*user.User{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"users"},
		dao.CreateColumnUser,
		r.toCopyFromSourceUsers(users...))
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) FindUser(email email.Email) (*user.User, error) {
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.findUserTx(ctx, tx, email)
}

func (r *Repository) findUserTx(ctx context.Context, tx pgx.Tx, email email.Email) (*user.User, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"updated_at",
		"email",
		"password",
		"name",
		"role",
		"token",
	).From("users")

	builder = builder.Where(squirrel.Eq{"email": email.String()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoUser []*dao.User
	if err = pgxscan.ScanAll(&daoUser, rows); err != nil {
		return nil, err
	}

	if len(daoUser) == 0 {
		return nil, errors.New("user not found")
	}

	return r.toDomainUser(daoUser[0])
}
