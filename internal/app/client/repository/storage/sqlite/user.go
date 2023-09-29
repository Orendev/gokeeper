package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/client/domain/user"
	"github.com/Orendev/gokeeper/internal/app/client/repository/storage/sqlite/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// AddUser create a user
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

// GetUser receive a user
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

// UpdateUser update user token
func (r *Repository) UpdateUser(ctx context.Context, id uuid.UUID, updateFn func(u *user.User) (*user.User, error)) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	upContact, err := r.oneUserTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	in, err := updateFn(upContact)
	if err != nil {
		return nil, err
	}

	return r.updateUserTx(ctx, tx, in)
}

func (r *Repository) updateUserTx(ctx context.Context, tx *sql.Tx, in *user.User) (*user.User, error) {

	builder := r.genSQL.
		Update("user").
		Set("password", in.Password().String()).
		Set("email", in.Email().String()).
		Set("role", in.Role().String()).
		Set("Name", in.Name().String()).
		Set("token", in.Token().String()).
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

func (r *Repository) oneUserTx(ctx context.Context, tx *sql.Tx, ID uuid.UUID) (*user.User, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"updated_at",
		"password",
		"email",
		"role",
		"name",
		"token",
	).From("user")

	builder = builder.Where(squirrel.Eq{"id": ID})

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
		return nil, err
	}

	return r.toDomainUser(daoUser[0])
}
