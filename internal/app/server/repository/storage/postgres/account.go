package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/account"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/pgxscan"
)

var mappingSortAccount = map[columnCode.ColumnCode]string{
	"id":      "id",
	"title":   "title",
	"version": "version",
}

func (r *Repository) CreateAccount(accounts ...*account.Account) ([]*account.Account, error) {
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createAccountTx(ctx, tx, accounts...)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createAccountTx(ctx context.Context, tx pgx.Tx, accounts ...*account.Account) ([]*account.Account, error) {
	if len(accounts) == 0 {
		return []*account.Account{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"keeper", "account"},
		dao.CreateColumnAccount,
		r.toCopyFromSourceAccounts(accounts...))
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *Repository) UpdateAccount(ID uuid.UUID, updateFn func(c *account.Account) (*account.Account, error)) (*account.Account, error) {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	upAccount, err := r.oneAccountTx(ctx, tx, ID)
	if err != nil {
		return nil, err
	}
	in, err := updateFn(upAccount)
	if err != nil {
		return nil, err
	}

	return r.updateAccountTx(ctx, tx, in)
}

func (r *Repository) updateAccountTx(ctx context.Context, tx pgx.Tx, in *account.Account) (*account.Account, error) {

	builder := r.genSQL.Update("keeper.account").
		Set("user_id", in.UserID()).
		Set("title", in.Title().String()).
		Set("login", in.Login().String()).
		Set("password", in.Password().String()).
		Set("updated_at", in.UpdatedAt()).
		Set("comment", in.Comment().String()).
		Set("web_address", in.WebAddress().String()).
		Set("version", in.Version()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":         in.ID(),
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
			web_address,
			version`,
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

	return r.toDomainAccount(daoAccounts[0])
}

func (r *Repository) DeleteAccount(ID uuid.UUID) error {
	var ctx = context.Background()

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
	builder := r.genSQL.Update("keeper.account").
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

func (r *Repository) ListAccount(parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	if parameter.Pagination.Limit == 0 {
		parameter.Pagination.Limit = r.options.DefaultLimit
	}

	accounts, err := r.listAccountTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *Repository) listAccountTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*account.Account, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"updated_at",
		"user_id",
		"title",
		"login",
		"password",
		"comment",
		"web_address",
		"version",
	).From("keeper.account")

	builder = builder.Where(squirrel.Eq{"is_deleted": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortAccount)...)
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

	return r.toDomainAccounts(daoAccounts)
}

func (r *Repository) ReadAccountByID(ID uuid.UUID) (response *account.Account, err error) {
	var ctx = context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err = r.oneAccountTx(ctx, tx, ID)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) oneAccountTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (*account.Account, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"updated_at",
		"user_id",
		"title",
		"login",
		"password",
		"comment",
		"web_address",
		"version",
	).From("keeper.account")

	builder = builder.Where(squirrel.Eq{"is_deleted": false, "id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoAccount []*dao.Account
	if err = pgxscan.ScanAll(&daoAccount, rows); err != nil {
		return nil, err
	}

	if len(daoAccount) == 0 {
		return nil, errors.New("Account not found")
	}

	return r.toDomainAccount(daoAccount[0])
}

func (r *Repository) CountAccount() (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From("keeper.account")

	builder = builder.Where(squirrel.Eq{"is_deleted": false})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var row = r.db.QueryRow(context.Background(), query, args...)
	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
