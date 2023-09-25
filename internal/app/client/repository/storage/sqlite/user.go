package sqlite

import (
	"context"
	"database/sql"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/google/uuid"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/pkg/errors"
)

func (r *Repository) AddUser(ctx context.Context, user user.User) (*user.User, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	response, err := r.addUserTx(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) addUserTx(ctx context.Context, tx *sql.Tx, user user.User) (*user.User, error) {

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

func (r *Repository) LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	panic("implement me")
}

func (r *Repository) GetUser(ctx context.Context) (*user.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.getUserTx(ctx, tx)
}

func (r *Repository) getUserTx(ctx context.Context, tx *sql.Tx) (*user.User, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"updated_at",
		"email",
		"role",
		"token",
		"password",
		"name",
	).From("user")

	builder = builder.Where(squirrel.NotEq{"email": nil})

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

func (r *Repository) UpdateToken(ctx context.Context, id uuid.UUID, token token.Token) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.updateTokenTx(ctx, tx, id, token)
}

func (r *Repository) updateTokenTx(ctx context.Context, tx *sql.Tx, id uuid.UUID, token token.Token) (*user.User, error) {

	builder := r.genSQL.Update("user").
		Set("token", token).
		Where(squirrel.And{
			squirrel.Eq{
				"id": id,
			},
		}).
		Suffix(`RETURNING
			id,
			created_at,
			updated_at,
			password,
			email,
			role,
			name,
			token`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoUser *dao.User
	if err = sqlscan.ScanOne(&daoUser, rows); err != nil {
		return nil, err
	}

	return r.toDomainUser(daoUser)
}