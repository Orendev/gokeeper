package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

var tableNameText = "texts"

func (r *Repository) CreateText(ctx context.Context, texts ...*text.TextData) ([]*text.TextData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	response, err := r.createTextTx(ctx, tx, texts...)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createTextTx(ctx context.Context, tx pgx.Tx, texts ...*text.TextData) ([]*text.TextData, error) {
	if len(texts) == 0 {
		return []*text.TextData{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{tableNameText},
		dao.CreateColumnData,
		r.toCopyFromSourceTexts(texts...))
	if err != nil {
		return nil, err
	}

	return texts, nil
}

func (r *Repository) UpdateText(ctx context.Context, id uuid.UUID, updateFn func(t *text.TextData) (*text.TextData, error)) (*text.TextData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	upText, err := r.oneTextTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	in, err := updateFn(upText)
	if err != nil {
		return nil, err
	}

	return r.updateTextTx(ctx, tx, in)
}

func (r *Repository) updateTextTx(ctx context.Context, tx pgx.Tx, in *text.TextData) (*text.TextData, error) {

	builder := r.genSQL.Update(tableNameText).
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

	var daoTexts []*dao.Data
	if err = pgxscan.ScanAll(&daoTexts, rows); err != nil {
		return nil, err
	}

	return r.toDomainText(daoTexts[0])
}

func (r *Repository) DeleteText(ctx context.Context, ID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteTextTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
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

	var daoTexts []*dao.Data
	if err = pgxscan.ScanAll(&daoTexts, rows); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListText(ctx context.Context, parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
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

	return texts, nil
}

func (r *Repository) listTextTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
	var builder = r.genSQL.Select(
		dao.CreateColumnData...,
	).From(tableNameText)

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

	var daoTexts []*dao.Data
	if err = pgxscan.ScanAll(&daoTexts, rows); err != nil {
		return nil, err
	}

	return r.toDomainTexts(daoTexts)
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

func (r *Repository) oneTextTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (*text.TextData, error) {
	var builder = r.genSQL.Select(
		dao.CreateColumnData...,
	).From(tableNameText)

	builder = builder.Where(squirrel.Eq{"is_deleted": false, "id": ID})

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

	if len(daoText) == 0 {
		return nil, ErrDataNotFound
	}

	return r.toDomainText(daoText[0])
}
