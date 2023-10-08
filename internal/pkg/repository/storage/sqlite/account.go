package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// CreateAccount create account
func (r *Repository) CreateAccount(ctx context.Context, account *account.Account) (*account.Account, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.createAccountTx(ctx, tx, account)
}

func (r *Repository) createAccountTx(ctx context.Context, tx *sql.Tx, account *account.Account) (*account.Account, error) {

	builder := r.genSQL.
		Insert(dao.TableNameAccount).
		Columns(dao.ColumnAccount...).
		Values(
			account.ID().String(),
			account.CreatedAt().String(),
			account.UpdatedAt().String(),
			account.UserID().String(),
			account.Password(),
			account.Login(),
			account.Title().String(),
			account.URL(),
			account.Comment(),
			account.IsDeleted(),
		).
		RunWith(tx)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// UpdateAccount update account
func (r *Repository) UpdateAccount(ctx context.Context, account *account.Account) (*account.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.updateAccountTx(ctx, tx, account)
}

func (r *Repository) updateAccountTx(ctx context.Context, tx *sql.Tx, account *account.Account) (*account.Account, error) {

	builder := r.genSQL.
		Update(dao.TableNameAccount).
		Set("password", account.Password()).
		Set("login", account.Login()).
		Set("title", account.Title().String()).
		Set("url", account.URL()).
		Set("comment", account.Comment()).
		Set("updated_at", time.Now().UTC()).
		Set("is_deleted", account.IsDeleted()).
		Where(squirrel.Eq{
			"id": account.ID().String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// ListAccount receive account
func (r *Repository) ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) (*account.ListAccountViewModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	accounts, err := r.listAccountTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountAccount(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &account.ListAccountViewModel{}

	list.Data = accounts
	list.Limit = parameter.Pagination.Limit
	list.Offset = parameter.Pagination.Offset
	list.Total = total

	return list, nil
}

func (r *Repository) listAccountTx(ctx context.Context, tx *sql.Tx, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
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

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoAccounts []*dao.Account
	if err = sqlscan.ScanAll(&daoAccounts, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainAccounts(daoAccounts)
}

// DeleteAccount delete account
func (r *Repository) DeleteAccount(ctx context.Context, id uuid.UUID) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.deleteAccountTx(ctx, tx, id)
}

func (r *Repository) deleteAccountTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	builder := r.genSQL.
		Update(dao.TableNameAccount).
		Set("updated_at", time.Now().UTC()).
		Set("is_deleted", true).
		Where(squirrel.Eq{
			"id": id,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	var daoAccounts []*dao.Account
	if err = sqlscan.ScanAll(&daoAccounts, rows); err != nil {
		return err
	}

	return nil
}

func (r Repository) CountAccount(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
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

	row := r.db.QueryRowContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
