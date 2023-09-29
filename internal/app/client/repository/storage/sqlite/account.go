package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/client/domain/account"
	"github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

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

// GetAccount receive account
func (r *Repository) GetAccount(ctx context.Context) (*account.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.getAccountTx(ctx, tx)
}

func (r *Repository) getAccountTx(ctx context.Context, tx *sql.Tx) (*account.Account, error) {
	var builder = r.genSQL.Select(
		dao.ColumnAccount...,
	).From("accounts")

	builder = builder.Where(squirrel.NotEq{"is_deleted": false})

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
