package card

import (
	"context"
	"os"
	"testing"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/mocks"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	storageRepository = new(mocks.Storage)
	ucCard            *UseCase
	data              = make(map[uuid.UUID]*card.CardData)
	cards             []*card.CardData
	listCardViewModel = &card.ListCardViewModel{}
	createCard        *card.CardData
	updateCard        *card.CardData
	parameter         = queryParameter.QueryParameter{}
	total             uint64
)

func TestMain(m *testing.M) {

	userID := uuid.New()

	createCard, _ = card.New(
		userID,
		[]byte("card number"),
		[]byte("card name"),
		[]byte("cvc"),
		[]byte("card date"),
		[]byte("comment"),
	)

	updateCard, _ = card.New(
		userID,
		[]byte("card number update"),
		[]byte("card name"),
		[]byte("cvc"),
		[]byte("card date"),
		[]byte("comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0
	total = 1

	os.Exit(m.Run())
}

func initTestUseCaseCard(t *testing.T) {
	assertion := assert.New(t)

	storageRepository.On("CreateCard",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, card *card.CardData) *card.CardData {
			assertion.Equal(card, createCard)
			cards = append(cards, card)
			listCardViewModel.Data = cards
			data[card.ID()] = card

			return card
		}, func(ctx context.Context, card *card.CardData) error {
			return nil
		})

	storageRepository.On("UpdateCard",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, card *card.CardData) *card.CardData {
			assertion.Equal(card, updateCard)
			data[card.ID()] = card
			listCardViewModel.Data = cards

			return card
		}, func(ctx context.Context, card *card.CardData) error {
			return nil
		})

	storageRepository.On("DeleteCard",
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

	storageRepository.On("ListCard",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) *card.ListCardViewModel {
			t := len(listCardViewModel.Data)
			listCardViewModel.Limit = parameter.Pagination.Limit
			listCardViewModel.Offset = parameter.Pagination.Offset
			listCardViewModel.Total = uint64(t)

			return listCardViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestCard(t *testing.T) {
	initTestUseCaseCard(t)
	ucCard = New(storageRepository, Options{})
	assertion := assert.New(t)

	t.Run("create card", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucCard.Create(ctx, createCard)
		assertion.NoError(err)
		assertion.Equal(result, createCard)
	})

	t.Run("update card", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucCard.Update(ctx, updateCard)
		assertion.NoError(err)
		assertion.Equal(result, updateCard)
	})

	t.Run("delete card", func(t *testing.T) {
		var ctx = context.Background()

		err := ucCard.Delete(ctx, createCard.ID())
		assertion.NoError(err)

	})

	t.Run("list card", func(t *testing.T) {
		var ctx = context.Background()
		total = uint64(len(listCardViewModel.Data))
		result, err := ucCard.List(ctx, parameter)
		assertion.NoError(err)
		assertion.Equal(result.Data, listCardViewModel.Data)
		assertion.Equal(result.Total, total)
	})
}
