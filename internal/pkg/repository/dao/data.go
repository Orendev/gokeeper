package dao

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
)

var TableNameText = "texts"
var TableNameBinary = "binaries"

// SortData Fields for sorting
var SortData = map[columnCode.ColumnCode]string{
	"id":    "id",
	"title": "title",
}

// Data description of fields in the database
type Data struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserId    uuid.UUID `db:"user_id"`
	Title     string    `db:"title"`
	Data      []byte    `db:"data"`
	Comment   []byte    `db:"comment"`
	IsDeleted bool      `db:"is_deleted"`
}

// ColumnData Names of fields in the database
var ColumnData = []string{
	"id",
	"created_at",
	"updated_at",
	"user_id",
	"title",
	"data",
	"comment",
	"is_deleted",
}

// ToDomainText Let's reduce the data to the structure text.TextData
func ToDomainText(dao *Data) (*text.TextData, error) {

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

// ToDomainTexts Let's reduce the data to the structure []text.TextData
func ToDomainTexts(dao []*Data) ([]*text.TextData, error) {
	var result = make([]*text.TextData, len(dao))
	for i, v := range dao {
		c, err := ToDomainText(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}

// ToDomainBinary Let's reduce the data to the structure binary.BinaryData
func ToDomainBinary(dao *Data) (*binary.BinaryData, error) {

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

// ToDomainBinaries Let's reduce the data to the structure []binary.BinaryData
func ToDomainBinaries(dao []*Data) ([]*binary.BinaryData, error) {
	var result = make([]*binary.BinaryData, len(dao))
	for i, v := range dao {
		c, err := ToDomainBinary(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
