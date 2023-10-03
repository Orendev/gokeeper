package data

import (
	"github.com/Orendev/gokeeper/internal/app/server/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/internal/pkg/repository/dao"
	"github.com/Orendev/gokeeper/pkg/type/columnCode"
	"github.com/Orendev/gokeeper/pkg/type/title"
)

// SortData Fields for sorting
var SortData = map[columnCode.ColumnCode]string{
	"id":    "id",
	"title": "title",
}

// ToDomainText Let's reduce the data to the structure text.TextData
func ToDomainText(dao *dao.Data) (*text.TextData, error) {

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
func ToDomainTexts(dao []*dao.Data) ([]*text.TextData, error) {
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
func ToDomainBinary(dao *dao.Data) (*binary.BinaryData, error) {

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

// ToDomainBinary Let's reduce the data to the structure []binary.BinaryData
func ToDomainBinaries(dao []*dao.Data) ([]*binary.BinaryData, error) {
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
