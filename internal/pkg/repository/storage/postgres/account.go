package postgres

import (
	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *Repository) CreateAccount(ctx context.Context, account *account.Account) (*account.Account, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createAccountTx(ctx, tx, account)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createAccountTx(ctx context.Context, tx pgx.Tx, account *account.Account) (*account.Account, error) {
	builder := r.genSQL.
		Insert(dao.TableNameAccount).
		Columns(dao.ColumnAccount...).
		Values(
			account.ID().String(),
			account.CreatedAt(),
			account.UpdatedAt(),
			account.UserID(),
			account.Password(),
			account.Login(),
			account.Title().String(),
			account.URL(),
			account.Comment(),
			account.IsDeleted(),
		).
		Suffix(`RETURNING
			id,
			created_at,
			updated_at,
			user_id,
			password,
			login,
			title,
			url,
			comment,
			is_deleted`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoUser []*dao.Account

	if err = pgxscan.ScanAll(&daoUser, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainAccount(daoUser[0])
}

// UpdateAccount update account
func (r *Repository) UpdateAccount(ctx context.Context, account *account.Account) (*account.Account, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.updateAccountTx(ctx, tx, account)
}

func (r *Repository) updateAccountTx(ctx context.Context, tx pgx.Tx, account *account.Account) (*account.Account, error) {

	builder := r.genSQL.Update("accounts").
		Set("user_id", account.UserID()).
		Set("title", account.Title().String()).
		Set("login", account.Login()).
		Set("password", account.Password()).
		Set("updated_at", account.UpdatedAt()).
		Set("comment", account.Comment()).
		Set("url", account.URL()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":         account.ID(),
				"is_deleted": false,
			},
		}).
		Suffix(`RETURNING
			id,
			created_at,
			updated_at,
			user_id,
			title,
			login,
			password,
			comment,
			url`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoAccounts []*dao.Account
	if err = pgxscan.ScanAll(&daoAccounts, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainAccount(daoAccounts[0])
}

func (r *Repository) DeleteAccount(ctx context.Context, ID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteAccountTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteAccountTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {
	builder := r.genSQL.Update("accounts").
		Set("is_deleted", true).
		Set("updated_at", time.Now().UTC()).
		Where(squirrel.Eq{"is_deleted": false, "id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	var daoAccounts []*dao.Account
	if err = pgxscan.ScanAll(&daoAccounts, rows); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) (*account.ListAccountViewModel, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	accounts, err := r.listAccountTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountBinary(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &account.ListAccountViewModel{
		Data:   accounts,
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  total,
	}

	return list, nil
}

func (r *Repository) listAccountTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	var builder = r.genSQL.Select(
		dao.ColumnAccount...,
	).From(dao.TableNameAccount)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	} else {
		builder = builder.Where(squirrel.Eq{
			"is_deleted": false,
		})
	}

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(dao.SortAccount)...)
	} else {
		builder = builder.OrderBy("created_at DESC")
	}

	if parameter.Pagination.Limit > 0 {
		builder = builder.Limit(parameter.Pagination.Limit)
	}
	if parameter.Pagination.Offset > 0 {
		builder = builder.Offset(parameter.Pagination.Offset)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoAccounts []*dao.Account
	if err = pgxscan.ScanAll(&daoAccounts, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainAccounts(daoAccounts)
}

func (r *Repository) CountAccount(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(dao.TableNameAccount)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	} else {
		builder = builder.Where(squirrel.Eq{
			"is_deleted": false,
		})
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
