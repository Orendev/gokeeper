package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/internal/pkg/repository/data"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// CreateBinary create binary data
func (r *Repository) CreateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.createBinaryTx(ctx, tx, binary)
}

func (r *Repository) createBinaryTx(ctx context.Context, tx *sql.Tx, binary *binary.BinaryData) (*binary.BinaryData, error) {

	builder := r.genSQL.
		Insert(dao.TableNameBinary).
		Columns(dao.ColumnCard...).
		Values(
			binary.ID().String(),
			binary.CreatedAt().String(),
			binary.UpdatedAt().String(),
			binary.UserID().String(),
			binary.Title().String(),
			binary.Data(),
			binary.Comment(),
			binary.IsDeleted(),
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

	return binary, nil
}

// UpdateBinary update binary data
func (r *Repository) UpdateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.updateBinaryTx(ctx, tx, binary)
}

func (r *Repository) updateBinaryTx(ctx context.Context, tx *sql.Tx, binary *binary.BinaryData) (*binary.BinaryData, error) {

	builder := r.genSQL.
		Update(dao.TableNameBinary).
		Set("user_id", binary.UserID()).
		Set("title", binary.Title().String()).
		Set("data", binary.Data()).
		Set("updated_at", binary.UpdatedAt()).
		Set("comment", binary.Comment()).
		Where(squirrel.Eq{
			"id": binary.ID().String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

// ListBinary receive binary data
func (r *Repository) ListBinary(ctx context.Context, parameter queryParameter.QueryParameter) (*binary.ListBinaryViewModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	binaries, err := r.listBinaryTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountBinary(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &binary.ListBinaryViewModel{}

	list.Data = binaries
	list.Limit = parameter.Pagination.Limit
	list.Offset = parameter.Pagination.Offset
	list.Total = total

	return list, nil
}

func (r *Repository) listBinaryTx(ctx context.Context, tx *sql.Tx, parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
	var builder = r.genSQL.Select(
		dao.ColumnData...,
	).From(dao.TableNameBinary)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	} else {
		builder = builder.Where(squirrel.Eq{
			"is_deleted": false,
		})
	}

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

	var daoBinaries []*dao.Data
	if err = sqlscan.ScanAll(&daoBinaries, rows); err != nil {
		return nil, err
	}

	return data.ToDomainBinaries(daoBinaries)
}

// DeleteBinary delete binary data
func (r *Repository) DeleteBinary(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteBinaryTx(ctx, tx, id); err != nil {
		return err
	}

	return nil

}

func (r *Repository) deleteBinaryTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	builder := r.genSQL.
		Update(dao.TableNameBinary).
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

	var daoBinaries []*dao.Data
	if err = sqlscan.ScanAll(&daoBinaries, rows); err != nil {
		return err
	}

	return nil
}

func (r Repository) CountBinary(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(dao.TableNameBinary)

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
