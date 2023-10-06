package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/internal/pkg/repository/data"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

// CreateBinary let's create a block of binary data
func (r *Repository) CreateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createBinaryTx(ctx, tx, binary)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createBinaryTx(ctx context.Context, tx pgx.Tx, binary *binary.BinaryData) (*binary.BinaryData, error) {
	builder := r.genSQL.
		Insert(dao.TableNameBinary).
		Columns(dao.ColumnData...).
		Values(
			binary.ID().String(),
			binary.CreatedAt(),
			binary.UpdatedAt(),
			binary.UserID().String(),
			binary.Title().String(),
			binary.Data(),
			binary.Comment(),
			binary.IsDeleted(),
		).
		Suffix(`RETURNING
			id,
			created_at,
			updated_at,
			user_id,
			title,
			data,
			comment`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoBinary []*dao.Data

	if err = pgxscan.ScanAll(&daoBinary, rows); err != nil {

		return nil, err
	}

	return data.ToDomainBinary(daoBinary[0])
}

func (r *Repository) UpdateBinary(ctx context.Context, binary *binary.BinaryData) (*binary.BinaryData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.updateBinaryTx(ctx, tx, binary)
}

func (r *Repository) updateBinaryTx(ctx context.Context, tx pgx.Tx, binary *binary.BinaryData) (*binary.BinaryData, error) {

	builder := r.genSQL.Update(dao.TableNameBinary).
		Set("title", binary.Title().String()).
		Set("data", binary.Data()).
		Set("updated_at", binary.UpdatedAt()).
		Set("comment", binary.Comment()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":         binary.ID(),
				"is_deleted": false,
			},
		}).
		Suffix(`RETURNING
			id,
			created_at,
			updated_at,
			user_id,
			title,
			data,
			comment`,
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var daoBinaries []*dao.Data
	if err = pgxscan.ScanAll(&daoBinaries, rows); err != nil {
		return nil, err
	}

	return data.ToDomainBinary(daoBinaries[0])
}

func (r *Repository) DeleteBinary(ctx context.Context, ID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteBinaryTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteBinaryTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {
	builder := r.genSQL.Update(dao.TableNameBinary).
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

	var daoBinaries []*dao.Data
	if err = pgxscan.ScanAll(&daoBinaries, rows); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListBinary(ctx context.Context, parameter queryParameter.QueryParameter) (*binary.ListBinaryViewModel, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	binaries, err := r.listBinaryTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountBinary(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &binary.ListBinaryViewModel{
		Data:   binaries,
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  total,
	}

	return list, nil
}

func (r *Repository) listBinaryTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
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
		builder = builder.OrderBy(parameter.Sorts.Parsing(data.SortData)...)
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

	var daoBinaries []*dao.Data
	if err = pgxscan.ScanAll(&daoBinaries, rows); err != nil {
		return nil, err
	}

	return data.ToDomainBinaries(daoBinaries)
}

func (r *Repository) CountBinary(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
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

	var row = r.db.QueryRow(ctx, query, args...)
	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
