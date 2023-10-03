package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/internal/pkg/repository/data"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// CreateText create text data
func (r *Repository) CreateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.createTextTx(ctx, tx, text)
}

func (r *Repository) createTextTx(ctx context.Context, tx *sql.Tx, text *text.TextData) (*text.TextData, error) {

	builder := r.genSQL.
		Insert(dao.TableNameText).
		Columns(dao.ColumnData...).
		Values(
			text.ID().String(),
			text.CreatedAt().String(),
			text.UpdatedAt().String(),
			text.UserID().String(),
			text.Title().String(),
			text.Data(),
			text.Comment(),
			text.IsDeleted(),
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

	return text, nil
}

// UpdateText update text data
func (r *Repository) UpdateText(ctx context.Context, text *text.TextData) (*text.TextData, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.updateTextTx(ctx, tx, text)
}

func (r *Repository) updateTextTx(ctx context.Context, tx *sql.Tx, text *text.TextData) (*text.TextData, error) {

	builder := r.genSQL.
		Update(dao.TableNameText).
		Set("user_id", text.UserID()).
		Set("title", text.Title().String()).
		Set("data", text.Data()).
		Set("updated_at", text.UpdatedAt()).
		Set("comment", text.Comment()).
		Where(squirrel.Eq{
			"id": text.ID().String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return text, nil
}

// ListText receive text
func (r *Repository) ListText(ctx context.Context, parameter queryParameter.QueryParameter) (*text.ListTextViewModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	texts, err := r.listTextTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountText(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &text.ListTextViewModel{}

	list.Data = texts
	list.Limit = parameter.Pagination.Limit
	list.Offset = parameter.Pagination.Offset
	list.Total = total

	return list, nil
}

func (r *Repository) listTextTx(ctx context.Context, tx *sql.Tx, parameter queryParameter.QueryParameter) ([]*text.TextData, error) {
	var builder = r.genSQL.Select(
		dao.ColumnData...,
	).From(dao.TableNameText)

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

	var daoTexts []*dao.Data
	if err = sqlscan.ScanAll(&daoTexts, rows); err != nil {
		return nil, err
	}

	return data.ToDomainTexts(daoTexts)
}

// DeleteText delete text data
func (r *Repository) DeleteText(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(r.options.Timeout)*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteTextTx(ctx, tx, id); err != nil {
		return err
	}

	return nil

}

func (r *Repository) deleteTextTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	builder := r.genSQL.
		Update(dao.TableNameText).
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

	var daoTexts []*dao.Data
	if err = sqlscan.ScanAll(&daoTexts, rows); err != nil {
		return err
	}

	return nil
}

func (r Repository) CountText(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(dao.TableNameText)

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
