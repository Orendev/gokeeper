package text

import (
	"context"
	"os"
	"testing"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/mocks"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	storageRepository = new(mocks.Storage)
	ucText            *UseCase
	data              = make(map[uuid.UUID]*text.TextData)
	texts             []*text.TextData
	listTextViewModel = &text.ListTextViewModel{}
	createText        *text.TextData
	updateText        *text.TextData
	parameter         = queryParameter.QueryParameter{}
	total             uint64
)

func TestMain(m *testing.M) {
	titleObj, _ := title.New("Test title")
	userID := uuid.New()

	createText, _ = text.New(
		userID,
		*titleObj,
		[]byte("text data"),
		[]byte("test comment"),
	)

	titleUpdateObj, _ := title.New("Test Update")

	updateText, _ = text.New(
		userID,
		*titleUpdateObj,
		[]byte("text data"),
		[]byte("test comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0
	total = 1

	os.Exit(m.Run())
}

func initTestUseCaseText(t *testing.T) {
	assertion := assert.New(t)

	storageRepository.On("CreateText",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, text *text.TextData) *text.TextData {
			assertion.Equal(text, createText)
			texts = append(texts, text)
			listTextViewModel.Data = texts
			data[text.ID()] = text

			return text
		}, func(ctx context.Context, text *text.TextData) error {
			return nil
		})

	storageRepository.On("UpdateText",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, text *text.TextData) *text.TextData {
			assertion.Equal(text, updateText)
			data[text.ID()] = text
			listTextViewModel.Data = texts

			return text
		}, func(ctx context.Context, text *text.TextData) error {
			return nil
		})

	storageRepository.On("DeleteText",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, id uuid.UUID) error {
			if _, ok := data[id]; !ok {
				return repository.ErrDataNotFound
			}
			return nil
		}, func(ctx context.Context, id uuid.UUID) error {
			return nil
		})

	storageRepository.On("ListText",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) *text.ListTextViewModel {
			t := len(listTextViewModel.Data)
			listTextViewModel.Limit = parameter.Pagination.Limit
			listTextViewModel.Offset = parameter.Pagination.Offset
			listTextViewModel.Total = uint64(t)

			return listTextViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestText(t *testing.T) {
	initTestUseCaseText(t)
	ucText = New(storageRepository, Options{})
	assertion := assert.New(t)

	t.Run("create text", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucText.Create(ctx, createText)
		assertion.NoError(err)
		assertion.Equal(result, createText)
	})

	t.Run("update text", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucText.Update(ctx, updateText)
		assertion.NoError(err)
		assertion.Equal(result, updateText)
	})

	t.Run("delete text", func(t *testing.T) {
		var ctx = context.Background()

		err := ucText.Delete(ctx, createText.ID())
		assertion.NoError(err)

	})

	t.Run("list text", func(t *testing.T) {
		var ctx = context.Background()
		total = uint64(len(listTextViewModel.Data))
		result, err := ucText.List(ctx, parameter)
		assertion.NoError(err)
		assertion.Equal(result.Data, listTextViewModel.Data)
		assertion.Equal(result.Total, total)
	})
}
