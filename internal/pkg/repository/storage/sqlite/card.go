package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// CreateCard create text data
func (r *Repository) CreateCard(ctx context.Context, card *card.CardData) (*card.CardData, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.createCardTx(ctx, tx, card)
}

func (r *Repository) createCardTx(ctx context.Context, tx *sql.Tx, card *card.CardData) (*card.CardData, error) {

	builder := r.genSQL.
		Insert(dao.TableNameCard).
		Columns(dao.ColumnCard...).
		Values(
			card.ID().String(),
			card.CreatedAt().String(),
			card.UpdatedAt().String(),
			card.UserID().String(),
			card.CardNumber(),
			card.CardName(),
			card.CardDate(),
			card.CVC(),
			card.Comment(),
			card.IsDeleted(),
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

	return card, nil
}

// UpdateCard update text data
func (r *Repository) UpdateCard(ctx context.Context, card *card.CardData) (*card.CardData, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	return r.updateCardTx(ctx, tx, card)
}

func (r *Repository) updateCardTx(ctx context.Context, tx *sql.Tx, card *card.CardData) (*card.CardData, error) {

	builder := r.genSQL.
		Update(dao.TableNameCard).
		Set("card_number", card.CardNumber()).
		Set("card_name", card.CardName()).
		Set("card_date", card.CardDate()).
		Set("cvv", card.CVC()).
		Set("comment", card.Comment()).
		Where(squirrel.Eq{
			"id": card.ID().String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// ListCard receive card data
func (r *Repository) ListCard(ctx context.Context, parameter queryParameter.QueryParameter) (*card.ListCardViewModel, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	cards, err := r.listCardTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountCard(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &card.ListCardViewModel{}

	list.Data = cards
	list.Limit = parameter.Pagination.Limit
	list.Offset = parameter.Pagination.Offset
	list.Total = total

	return list, nil
}

func (r *Repository) listCardTx(ctx context.Context, tx *sql.Tx, parameter queryParameter.QueryParameter) ([]*card.CardData, error) {
	var builder = r.genSQL.Select(
		dao.ColumnCard...,
	).From(dao.TableNameCard)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	} else {
		builder = builder.Where(squirrel.Eq{
			"is_deleted": false,
		})
	}

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(dao.SortCard)...)
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

	var daoCards []*dao.Card
	if err = sqlscan.ScanAll(&daoCards, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainCards(daoCards)
}

// DeleteCard delete text data
func (r *Repository) DeleteCard(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t *sql.Tx) {
		err = transaction.FinishSQL(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteCardTx(ctx, tx, id); err != nil {
		return err
	}

	return nil

}

func (r *Repository) deleteCardTx(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	builder := r.genSQL.
		Update(dao.TableNameCard).
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

	var daoCards []*dao.Card
	if err = sqlscan.ScanAll(&daoCards, rows); err != nil {
		return err
	}

	return nil
}

func (r Repository) CountCard(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From(dao.TableNameCard)

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
