package sqlite

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/hashedPassword"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/sqlscan"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUser create a user
func (r *Repository) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.createUserTx(ctx, tx, user)
}

func (r *Repository) createUserTx(ctx context.Context, tx *sql.Tx, user *user.User) (*user.User, error) {
	hashPasswordUser, err := hashedPassword.New(user.Password().String())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	passwordUser, err := password.New(hashPasswordUser.String())
	if err != nil {
		logger.Log.Error("error create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "user password validation error: %v", err)
	}

	builder := r.genSQL.
		Insert(dao.TableNameUser).
		Columns(dao.ColumnUser...).
		Values(
			user.ID().String(),
			passwordUser,
			user.Email().String(),
			user.Name().String(),
			user.Role().String(),
			user.Token().String(),
			user.CreatedAt(),
			user.UpdatedAt(),
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

	return user, nil
}

// LoginUser receive a user
func (r *Repository) LoginUser(ctx context.Context, email email.Email, password password.Password) (*user.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
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

func (r *Repository) findUserTx(ctx context.Context, tx *sql.Tx, email email.Email) (*user.User, error) {
	var builder = r.genSQL.Select(
		dao.ColumnUser...,
	).From(dao.TableNameUser)

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
		return nil, repository.ErrNotFoundUser
	}

	return dao.ToDomainUser(daoUser[0])
}

func (r *Repository) SetTokenUser(ctx context.Context, user *user.User) bool {
	tx, err := r.db.Begin()
	if err != nil {
		return false
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.updateUserTokenTx(ctx, tx, user)
}

func (r *Repository) updateUserTokenTx(ctx context.Context, tx *sql.Tx, user *user.User) bool {

	builder := r.genSQL.Update(dao.TableNameUser).
		Set("token", user.Token()).
		Where(squirrel.Eq{
			"id": user.ID().String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return false
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return false
	}

	total, err := res.RowsAffected()
	if err != nil {
		return false
	}
	if total == 0 {
		return false
	}

	return true
}

func (r Repository) CountUser(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(dao.TableNameUser)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
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
