package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/tools/transaction"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
)

func (r *Repository) CreateCard(ctx context.Context, card *card.CardData) (*card.CardData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.createCardTx(ctx, tx, card)
}

func (r *Repository) createCardTx(ctx context.Context, tx pgx.Tx, card *card.CardData) (*card.CardData, error) {
	builder := r.genSQL.
		Insert(dao.TableNameCard).
		Columns(dao.ColumnCard...).
		Values(
			card.ID().String(),
			card.CreatedAt(),
			card.UpdatedAt(),
			card.UserID().String(),
			card.CardNumber(),
			card.CardName(),
			card.CardDate(),
			card.CVC(),
			card.Comment(),
		).
		Suffix(`RETURNING
				id,
			created_at,
			updated_at,
			user_id,
			card_number,
			card_name,
			card_date,
			cvv,
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

	var daoCard []*dao.Card

	if err = pgxscan.ScanAll(&daoCard, rows); err != nil {

		return nil, err
	}

	return dao.ToDomainCard(daoCard[0])
}

func (r *Repository) UpdateCard(ctx context.Context, card *card.CardData) (*card.CardData, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	return r.updateCardTx(ctx, tx, card)
}

func (r *Repository) updateCardTx(ctx context.Context, tx pgx.Tx, card *card.CardData) (*card.CardData, error) {

	builder := r.genSQL.Update(dao.TableNameCard).
		Set("card_number", card.CardNumber()).
		Set("card_name", card.CardName()).
		Set("card_date", card.CardDate()).
		Set("cvv", card.CVC()).
		Set("comment", card.Comment()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":         card.ID(),
				"is_deleted": false,
			},
		}).
		Suffix(`RETURNING
			id,
			created_at,
			updated_at,
			user_id,
			card_number,
			card_name,
			card_date,
			cvv,
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

	var daoCards []*dao.Card
	if err = pgxscan.ScanAll(&daoCards, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainCard(daoCards[0])
}

func (r *Repository) DeleteCard(ctx context.Context, ID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteCardTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteCardTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {
	builder := r.genSQL.Update(dao.TableNameCard).
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

	var daoCards []*dao.Card
	if err = pgxscan.ScanAll(&daoCards, rows); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListCard(ctx context.Context, parameter queryParameter.QueryParameter) (*card.ListCardViewModel, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.FinishPGX(ctx, t, err)
	}(ctx, tx)

	cards, err := r.listCardTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	total, err := r.CountCard(ctx, parameter)
	if err != nil {
		return nil, err
	}

	list := &card.ListCardViewModel{
		Data:   cards,
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  total,
	}

	return list, nil
}

func (r *Repository) listCardTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*card.CardData, error) {
	var builder = r.genSQL.Select(
		dao.ColumnData...,
	).From(dao.TableNameCard)

	if len(parameter.Filters) > 0 {
		builder = builder.Where(parameter.Filters.Eq())
	} else {
		builder = builder.Where(squirrel.Eq{
			"is_deleted": false,
		})
	}

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(dao.SortData)...)
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

	var daoCards []*dao.Card
	if err = pgxscan.ScanAll(daoCards, rows); err != nil {
		return nil, err
	}

	return dao.ToDomainCards(daoCards)
}

func (r *Repository) CountCard(ctx context.Context, parameter queryParameter.QueryParameter) (uint64, error) {
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

	var row = r.db.QueryRow(ctx, query, args...)
	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
