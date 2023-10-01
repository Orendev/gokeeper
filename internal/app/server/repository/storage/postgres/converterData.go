package postgres

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/internal/app/server/domain/text"
	"github.com/Orendev/gokeeper/internal/app/server/repository/storage/postgres/dao"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/jackc/pgx/v5"
)

var mappingSortData = map[columnCode.ColumnCode]string{
	"id":    "id",
	"title": "title",
}

func (r Repository) toCopyFromSourceTexts(texts ...*text.TextData) pgx.CopyFromSource {
	rows := make([][]interface{}, len(texts))

	for i, val := range texts {

		rows[i] = []interface{}{
			val.ID().String(),
			val.CreatedAt().UTC(),
			val.UpdatedAt().UTC(),
			val.UserID().String(),
			val.Title().String(),
			val.Data(),
			val.Comment(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainText(dao *dao.Data) (*text.TextData, error) {

	titleObject, err := title.New(dao.Title)
	if err != nil {
		return nil, err
	}

	return text.NewWithID(
		dao.ID,
		dao.UserId,
		*titleObject,
		dao.Data,
		dao.Comment,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
}

func (r Repository) toDomainTexts(dao []*dao.Data) ([]*text.TextData, error) {
	var result = make([]*text.TextData, len(dao))
	for i, v := range dao {
		c, err := r.toDomainText(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}

func (r Repository) toCopyFromSourceBinary(binaries ...*binary.BinaryData) pgx.CopyFromSource {
	rows := make([][]interface{}, len(binaries))

	for i, val := range binaries {

		rows[i] = []interface{}{
			val.ID().String(),
			val.CreatedAt().UTC(),
			val.UpdatedAt().UTC(),
			val.UserID().String(),
			val.Title().String(),
			val.Data(),
			val.Comment(),
		}
	}
	// Use CopyFrom to efficiently insert multiple rows at a time using the PostgreSQL copy protocol
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainBinary(dao *dao.Data) (*binary.BinaryData, error) {

	titleObject, err := title.New(dao.Title)
	if err != nil {
		return nil, err
	}

	return binary.NewWithID(
		dao.ID,
		dao.UserId,
		*titleObject,
		dao.Data,
		dao.Comment,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
}

func (r Repository) toDomainBinaries(dao []*dao.Data) ([]*binary.BinaryData, error) {
	var result = make([]*binary.BinaryData, len(dao))
	for i, v := range dao {
		c, err := r.toDomainBinary(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
