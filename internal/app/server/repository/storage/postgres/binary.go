package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

var tableNameBinary = "binaries"

func (r *Repository) CreateBinary(ctx context.Context, binaries ...*binary.BinaryData) ([]*binary.BinaryData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createBinaryTx(ctx, tx, binaries...)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createBinaryTx(ctx context.Context, tx pgx.Tx, binaries ...*binary.BinaryData) ([]*binary.BinaryData, error) {
	if len(binaries) == 0 {
		return []*binary.BinaryData{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{tableNameBinary},
		dao.CreateColumnData,
		r.toCopyFromSourceBinary(binaries...))
	if err != nil {
		return nil, err
	}

	return binaries, nil
}

func (r *Repository) UpdateBinary(ctx context.Context, id uuid.UUID, updateFn func(t *binary.BinaryData) (*binary.BinaryData, error)) (*binary.BinaryData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	upBinary, err := r.oneBinaryTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	in, err := updateFn(upBinary)
	if err != nil {
		return nil, err
	}

	return r.updateBinaryTx(ctx, tx, in)
}

func (r *Repository) updateBinaryTx(ctx context.Context, tx pgx.Tx, in *binary.BinaryData) (*binary.BinaryData, error) {

	builder := r.genSQL.Update(tableNameBinary).
		Set("user_id", in.UserID()).
		Set("title", in.Title().String()).
		Set("data", in.Data()).
		Set("updated_at", in.UpdatedAt()).
		Set("comment", in.Comment()).
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

	return r.toDomainBinary(daoBinaries[0])
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
	builder := r.genSQL.Update(tableNameBinary).
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

func (r *Repository) ListBinary(ctx context.Context, parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	Binaries, err := r.listBinaryTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	return Binaries, nil
}

func (r *Repository) listBinaryTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*binary.BinaryData, error) {
	var builder = r.genSQL.Select(
		dao.CreateColumnData...,
	).From(tableNameBinary)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	} else {
		builder = builder.Where(squirrel.Eq{
			"is_deleted": false,
		})
	}

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortData)...)
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

	return r.toDomainBinaries(daoBinaries)
}

func (r *Repository) CountBinary(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(tableNameBinary)

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

func (r *Repository) oneBinaryTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (*binary.BinaryData, error) {
	var builder = r.genSQL.Select(
		dao.CreateColumnData...,
	).From(tableNameBinary)

	builder = builder.Where(squirrel.Eq{"is_deleted": false, "id": ID})

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

	if len(daoBinary) == 0 {
		return nil, ErrDataNotFound
	}

	return r.toDomainBinary(daoBinary[0])
}
