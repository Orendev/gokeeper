package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *Repository) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createUserTx(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createUserTx(ctx context.Context, tx pgx.Tx, user *user.User) (*user.User, error) {
	builder := r.genSQL.
		Insert(dao.TableNameUser).
		Columns(dao.ColumnUser...).
		Values(
			user.ID().String(),
			user.Password().String(),
			user.Email().String(),
			user.Name().String(),
			user.Role().String(),
			user.Token().String(),
			user.CreatedAt(),
			user.UpdatedAt(),
		).
		Suffix(`RETURNING
			id,
			password,
			email,
			name,
			role,
			token,
			created_at,
			updated_at`,
		)

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

	return dao.ToDomainUser(daoUser[0])

}

func (r *Repository) SetTokenUser(ctx context.Context, user *user.User) bool {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return false
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.updateUserTokenTx(ctx, tx, user)
}

func (r *Repository) updateUserTokenTx(ctx context.Context, tx pgx.Tx, user *user.User) bool {

	builder := r.genSQL.Update(dao.TableNameUser).
		Set("token", user.Token()).
		Where(squirrel.And{
			squirrel.Eq{
				"id": user.ID().String(),
			},
		}).
		Suffix(`RETURNING
			id,
			password,
			email,
			name,
			role,
			token,
			created_at,
			updated_at`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return false
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return false
	}

	var daoUsers []*dao.User
	if err = pgxscan.ScanAll(&daoUsers, rows); err != nil {
		return false
	}

	return true
}

func (r *Repository) LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	u, err := r.findUserTx(ctx, tx, email)
	if err != nil {
		return nil, err
	}

	if !u.IsCorrectPassword(password) {
		return nil, repository.ErrIncorrectPassword
	}

	return u, nil
}

func (r *Repository) findUserTx(ctx context.Context, tx pgx.Tx, email email.Email) (*user.User, error) {
	var builder = r.genSQL.Select(
		dao.ColumnUser...,
	).From(dao.TableNameUser)

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

	return dao.ToDomainUser(daoUser[0])
}

func (r *Repository) CountUser(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(tableNameText)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var row = r.db.QueryRow(ctx, query, args...)
	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
