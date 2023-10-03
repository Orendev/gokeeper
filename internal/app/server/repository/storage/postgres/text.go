package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/internal/pkg/repository/data"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

var tableNameText = "texts"

func (r *Repository) CreateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.createTextTx(ctx, tx, text)
}

func (r *Repository) createTextTx(ctx context.Context, tx pgx.Tx, text *text.TextData) (*text.TextData, error) {

	builder := r.genSQL.
		Insert(dao.TableNameText).
		Columns(dao.ColumnData...).
		Values(
			text.ID().String(),
			text.CreatedAt(),
			text.UpdatedAt(),
			text.UserID().String(),
			text.Title().String(),
			text.Data(),
			text.Comment(),
			text.IsDeleted(),
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

	var daoText []*dao.Data

	if err = pgxscan.ScanAll(&daoText, rows); err != nil {

		return nil, err
	}

	return data.ToDomainText(daoText[0])
}

func (r *Repository) UpdateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.updateTextTx(ctx, tx, text)
}

func (r *Repository) updateTextTx(ctx context.Context, tx pgx.Tx, text *text.TextData) (*text.TextData, error) {

	builder := r.genSQL.Update(tableNameText).
		Set("title", text.Title().String()).
		Set("data", text.Data()).
		Set("updated_at", text.UpdatedAt()).
		Set("comment", text.Comment()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":         text.ID(),
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
	var daoText []*dao.Data

	if err = pgxscan.ScanAll(&daoText, rows); err != nil {
		return nil, err
	}

	return data.ToDomainText(daoText[0])
}

func (r *Repository) DeleteText(ctx context.Context, ID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.deleteTextTx(ctx, tx, ID)
}

func (r *Repository) deleteTextTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {
	builder := r.genSQL.Update(tableNameText).
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

	var daoText []*dao.Data

	return pgxscan.ScanAll(&daoText, rows)
}

func (r *Repository) ListText(ctx context.Context, parameter queryParameter.QueryParameter) (*text.ListTextViewModel, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	texts, err := r.listTextTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountText(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &text.ListTextViewModel{
		Data:   texts,
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  total,
	}

	return list, nil
}

func (r *Repository) listTextTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
	var builder = r.genSQL.Select(
		dao.ColumnData...,
	).From(tableNameText)

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

	var daoTexts []*dao.Data
	if err = pgxscan.ScanAll(&daoTexts, rows); err != nil {
		return nil, err
	}

	return data.ToDomainTexts(daoTexts)
}

func (r *Repository) CountText(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(tableNameText)

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
