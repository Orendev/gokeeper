package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var mappingSortContact = map[columnCode.ColumnCode]string{
	"id":    "id",
	"title": "title",
}

// CreateAccount create account
func (r *Repository) CreateAccount(ctx context.Context, account account.Account) (*account.Account, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	res, err := r.createAccountTx(ctx, tx, account)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) createAccountTx(ctx context.Context, tx *sql.Tx, account account.Account) (*account.Account, error) {

	builder := r.genSQL.
		Insert("accounts").
		Columns(dao.ColumnAccount...).
		Values(
			account.ID().String(),
			account.CreatedAt().String(),
			account.UpdatedAt().String(),
			account.Password().String(),
			account.Login().String(),
			account.Title().String(),
			account.URL().String(),
			account.Comment().String(),
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

	return &account, nil
}

// GetByIDAccount receive account
func (r *Repository) GetByIDAccount(ctx context.Context, id uuid.UUID) (*account.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.getAccountTx(ctx, tx, id)
}

func (r *Repository) getAccountTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*account.Account, error) {
	var builder = r.genSQL.Select(
		dao.ColumnAccount...,
	).From("accounts")

	builder = builder.Where(squirrel.NotEq{"is_deleted": false, "id": id})

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

	if len(daoAccounts) == 0 {
		return nil, errors.New("account not found")
	}

	return r.toDomainAccount(daoAccounts[0])
}

// UpdateAccount update account
func (r *Repository) UpdateAccount(ctx context.Context, id uuid.UUID, updateFn func(update *account.Account) (*account.Account, error)) (*account.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	upAccount, err := r.oneAccountTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	in, err := updateFn(upAccount)
	if err != nil {
		return nil, err
	}

	return r.updateAccountTx(ctx, tx, in)
}

func (r *Repository) updateAccountTx(ctx context.Context, tx *sql.Tx, in *account.Account) (*account.Account, error) {

	builder := r.genSQL.
		Update("accounts").
		Set("password", in.Password().String()).
		Set("login", in.Login().String()).
		Set("title", in.Title().String()).
		Set("url", in.URL().String()).
		Set("comment", in.Comment().String()).
		Set("updated_at", time.Now().UTC()).
		Set("is_deleted", in.IsDeleted()).
		Where(squirrel.Eq{
			"id": in.ID().String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (r *Repository) oneAccountTx(ctx context.Context, tx *sql.Tx, ID uuid.UUID) (*account.Account, error) {
	var builder = r.genSQL.Select(
		dao.ColumnAccount...,
	).From("accounts")

	builder = builder.Where(squirrel.Eq{"id": ID})

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

	if len(daoAccounts) == 0 {
		return nil, err
	}

	return r.toDomainAccount(daoAccounts[0])
}

// ListAccount receive account
func (r *Repository) ListAccount(ctx context.Context, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.listAccountTx(ctx, tx, parameter)
}

func (r *Repository) listAccountTx(ctx context.Context, tx *sql.Tx, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	var builder = r.genSQL.Select(
		dao.ColumnAccount...,
	).From("accounts")

	builder = builder.Where(squirrel.Eq{"is_deleted": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortContact)...)
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

	return r.toDomainAccounts(daoAccounts)
}

// DeleteAccount delete account
func (r *Repository) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteAccountTx(ctx, tx, id); err != nil {
		return err
	}

	return nil

}

func (r *Repository) deleteAccountTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	builder := r.genSQL.
		Update("accounts").
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
