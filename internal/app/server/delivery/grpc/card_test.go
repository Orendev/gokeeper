package grpc

import (
	"fmt"
	"testing"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/card"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	useCaseCard "github.com/Orendev/gokeeper/internal/pkg/useCase/card"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

var (
	cards             []*card.CardData
	dataCard          = make(map[uuid.UUID]*card.CardData)
	listCardViewModel = &card.ListCardViewModel{}
	createCard        *card.CardData
	updateCard        *card.CardData
	totalCard         uint64
	createReqCard     *protobuff.CreateCardRequest
	createResCard     *protobuff.CreateCardResponse

	updateReqCard *protobuff.UpdateCardRequest
	updateResCard *protobuff.UpdateCardResponse

	deleteReqCard *protobuff.DeleteCardRequest
	deleteResCard *protobuff.DeleteCardResponse

	listReqCard *protobuff.ListCardRequest
	listResCard *protobuff.ListCardResponse
)

func testMainCard() {
	createCard, _ = card.New(
		userID,
		[]byte("9836-9239-5555-1338"),
		[]byte("Ivanov Ivan"),
		[]byte("568"),
		[]byte("02/23"),
		[]byte("test comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	createCardProto := &protobuff.Card{
		ID:         createCard.ID().String(),
		CardNumber: createCard.CardNumber(),
		CardName:   createCard.CardName(),
		CVC:        createCard.CVC(),
		CardDate:   createCard.CardDate(),
		Comment:    createCard.Comment(),
		CreatedAt:  createCard.CreatedAt().Format(time.RFC3339),
		UpdatedAt:  createCard.UpdatedAt().Format(time.RFC3339),
		UserID:     createCard.UserID().String(),
	}
	createReqCard = &protobuff.CreateCardRequest{Data: createCardProto}

	createResCard = &protobuff.CreateCardResponse{
		ID: createCard.ID().String(),
	}

	updateCard, _ = card.NewWithID(
		createCard.ID(),
		userID,
		[]byte("9999-9239-5555-1338"),
		[]byte("Ivanov Ivan"),
		[]byte("568"),
		[]byte("02/25"),
		[]byte("test comment update"),
		createCard.CreatedAt().UTC(),
		time.Now().UTC(),
	)

	updateCardProto := &protobuff.Card{
		ID:         updateCard.ID().String(),
		CardNumber: updateCard.CardNumber(),
		CardName:   updateCard.CardName(),
		CVC:        updateCard.CVC(),
		CardDate:   updateCard.CardDate(),
		Comment:    updateCard.Comment(),
		CreatedAt:  updateCard.CreatedAt().Format(time.RFC3339),
		UpdatedAt:  updateCard.UpdatedAt().Format(time.RFC3339),
		UserID:     updateCard.UserID().String(),
	}

	updateReqCard = &protobuff.UpdateCardRequest{Data: updateCardProto}

	updateResCard = &protobuff.UpdateCardResponse{
		ID: updateCard.ID().String(),
	}

	deleteReqCard = &protobuff.DeleteCardRequest{
		ID: updateCard.ID().String(),
	}
	deleteResCard = &protobuff.DeleteCardResponse{
		ID: updateCard.ID().String(),
	}

	listReqCard = &protobuff.ListCardRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	totalCard = 1

	listCardProto := []*protobuff.Card{}

	listCardProto = append(listCardProto, updateCardProto)
	listResCard = &protobuff.ListCardResponse{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  totalCard,
		Data:   listCardProto,
	}
}

func initTestUseCaseCard(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateCard",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, card *card.CardData) *card.CardData {
			assertion.Equal(card.ID(), updateCard.ID())
			cards = append(cards, card)
			listCardViewModel.Data = cards
			dataCard[card.ID()] = card

			return card
		}, func(ctx context.Context, card *card.CardData) error {
			return nil
		})

	storageRepository.On("UpdateCard",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, card *card.CardData) *card.CardData {
			assertion.Equal(card.ID(), updateCard.ID())
			if len(cards) > 0 {
				cards[0] = card
			} else {
				cards = append(cards, card)
			}

			dataCard[card.ID()] = card
			listCardViewModel.Data = cards

			return card
		}, func(ctx context.Context, card *card.CardData) error {
			return nil
		})

	storageRepository.On("DeleteCard",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, id uuid.UUID) error {
			if _, ok := dataCard[id]; !ok {
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

			listCardViewModel.Limit = parameter.Pagination.Limit
			listCardViewModel.Offset = parameter.Pagination.Offset
			listCardViewModel.Total = uint64(len(listCardViewModel.Data))

			return listCardViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestDeliveryCard(t *testing.T) {
	initTestUseCaseCard(t)
	ucCard = useCaseCard.New(storageRepository, useCaseCard.Options{})
	option := Options{}
	deliveryGRPC := New(ucUser, ucAccount, ucCard, ucText, ucBinary, jwtManager, option)
	assertion := assert.New(t)

	t.Run("positive create card", func(t *testing.T) {
		result, err := deliveryGRPC.CreateCard(ctx, createReqCard)
		assertion.NoError(err)
		assertion.Equal(result, createResCard)
	})

	t.Run("negative create card", func(t *testing.T) {
		result, err := deliveryGRPC.CreateCard(context.Background(), createReqCard)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive update card", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateCard(ctx, updateReqCard)
		assertion.NoError(err)
		assertion.Equal(result, updateResCard)
	})

	t.Run("negative update card", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateCard(context.Background(), updateReqCard)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive delete card", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteCard(ctx, deleteReqCard)
		assertion.NoError(err)
		assertion.Equal(result, deleteResCard)
	})

	t.Run("negative delete card", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteCard(context.Background(), &protobuff.DeleteCardRequest{ID: uuid.New().String()})
		assertion.Errorf(err, "")
		assertion.Nil(result)
	})

	t.Run("positive list card", func(t *testing.T) {
		totalCard = uint64(len(listCardViewModel.Data))
		result, err := deliveryGRPC.ListCard(ctx, listReqCard)
		assertion.NoError(err)
		fmt.Println("tot", totalCard)
		assertion.Equal(result.Data[0], listResCard.Data[0])
		assertion.Equal(result.Total, totalCard)
	})

	t.Run("negative list card", func(t *testing.T) {
		result, err := deliveryGRPC.ListCard(context.Background(), listReqCard)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})
}
