package grpc

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/text"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/mocks"
	"github.com/Orendev/gokeeper/internal/pkg/useCase"
	useCaseText "github.com/Orendev/gokeeper/internal/pkg/useCase/text"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

var (
	storageRepository = new(mocks.Storage)
	// text
	texts             []*text.TextData
	dataText          = make(map[uuid.UUID]*text.TextData)
	listTextViewModel = &text.ListTextViewModel{}
	createText        *text.TextData
	updateText        *text.TextData
	parameter         = queryParameter.QueryParameter{}
	totalText         uint64

	createReqText *protobuff.CreateTextRequest
	createResText *protobuff.CreateTextResponse

	updateReqText *protobuff.UpdateTextRequest
	updateResText *protobuff.UpdateTextResponse

	deleteReqText *protobuff.DeleteTextRequest
	deleteResText *protobuff.DeleteTextResponse

	listReqText *protobuff.ListTextRequest
	listResText *protobuff.ListTextResponse

	ucUser     useCase.User
	ucAccount  useCase.Account
	ucCard     useCase.Card
	ucText     useCase.Text
	ucBinary   useCase.Binary
	jwtManager *auth.JWTManager
	ctx        context.Context
	userID     uuid.UUID
)

func TestMain(m *testing.M) {

	testMainText()
	testMainBinary()
	testMainCard()

	os.Exit(m.Run())
}

func testMainText() {
	jwtManager = auth.NewJWTManager("", 120*time.Second)
	userID = uuid.New()
	accessToken, _ := jwtManager.Generate(userID)
	header := metadata.New(map[string]string{auth.AuthorizationKey: accessToken})
	ctx = metadata.NewIncomingContext(context.Background(), header)

	titleObj, _ := title.New("Test title")

	createText, _ = text.New(
		userID,
		*titleObj,
		[]byte("test ya.ru"),
		[]byte("test comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	txData := &protobuff.Data{
		ID:        createText.ID().String(),
		Title:     createText.Title().String(),
		Data:      createText.Data(),
		Comment:   createText.Comment(),
		CreatedAt: createText.CreatedAt().Format(time.RFC3339),
		UpdatedAt: createText.UpdatedAt().Format(time.RFC3339),
		UserID:    createText.UserID().String(),
	}
	createReqText = &protobuff.CreateTextRequest{Data: txData}

	createResText = &protobuff.CreateTextResponse{
		ID: createText.ID().String(),
	}

	titleUpdateObj, _ := title.New("Test Update")

	updateText, _ = text.NewWithID(
		createText.ID(),
		userID,
		*titleUpdateObj,
		[]byte("text data update"),
		[]byte("test comment update"),
		createText.CreatedAt().UTC(),
		time.Now().UTC(),
	)

	txUpdateData := &protobuff.Data{
		ID:        updateText.ID().String(),
		Title:     updateText.Title().String(),
		Data:      updateText.Data(),
		Comment:   updateText.Comment(),
		CreatedAt: updateText.CreatedAt().Format(time.RFC3339),
		UpdatedAt: updateText.UpdatedAt().Format(time.RFC3339),
		UserID:    updateText.UserID().String(),
	}

	updateReqText = &protobuff.UpdateTextRequest{Data: txUpdateData}

	updateResText = &protobuff.UpdateTextResponse{
		ID: updateText.ID().String(),
	}

	deleteReqText = &protobuff.DeleteTextRequest{
		ID: updateText.ID().String(),
	}
	deleteResText = &protobuff.DeleteTextResponse{
		ID: updateText.ID().String(),
	}

	listReqText = &protobuff.ListTextRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	totalText = 1

	txListData := []*protobuff.Data{}

	txListData = append(txListData, txUpdateData)
	listResText = &protobuff.ListTextResponse{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
		Total:  totalText,
		Data:   txListData,
	}
}

func initTestUseCaseText(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateText",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, text *text.TextData) *text.TextData {
			assertion.Equal(text.ID(), updateText.ID())
			texts = append(texts, text)
			listTextViewModel.Data = texts
			dataText[text.ID()] = text

			return text
		}, func(ctx context.Context, text *text.TextData) error {
			return nil
		})

	storageRepository.On("UpdateText",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, text *text.TextData) *text.TextData {
			assertion.Equal(text.ID(), updateText.ID())
			if len(texts) > 0 {
				texts[0] = text
			} else {
				texts = append(texts, text)
			}

			dataText[text.ID()] = text
			listTextViewModel.Data = texts

			return text
		}, func(ctx context.Context, text *text.TextData) error {
			return nil
		})

	storageRepository.On("DeleteText",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, id uuid.UUID) error {
			if _, ok := dataText[id]; !ok {
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

			listTextViewModel.Limit = parameter.Pagination.Limit
			listTextViewModel.Offset = parameter.Pagination.Offset
			listTextViewModel.Total = uint64(len(listTextViewModel.Data))

			return listTextViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestDeliveryText(t *testing.T) {
	initTestUseCaseText(t)
	ucText = useCaseText.New(storageRepository, useCaseText.Options{})
	option := Options{}
	deliveryGRPC := New(ucUser, ucAccount, ucCard, ucText, ucBinary, jwtManager, option)
	assertion := assert.New(t)

	t.Run("positive create text", func(t *testing.T) {
		result, err := deliveryGRPC.CreateText(ctx, createReqText)
		assertion.NoError(err)
		assertion.Equal(result, createResText)
	})

	t.Run("negative create text", func(t *testing.T) {
		result, err := deliveryGRPC.CreateText(context.Background(), createReqText)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive update text", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateText(ctx, updateReqText)
		assertion.NoError(err)
		assertion.Equal(result, updateResText)
	})

	t.Run("negative update text", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateText(context.Background(), updateReqText)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive delete text", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteText(ctx, deleteReqText)
		assertion.NoError(err)
		assertion.Equal(result, deleteResText)
	})

	t.Run("negative delete text", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteText(context.Background(), &protobuff.DeleteTextRequest{ID: uuid.New().String()})
		assertion.Errorf(err, "")
		assertion.Nil(result)
	})

	t.Run("positive list text", func(t *testing.T) {
		totalText = uint64(len(listTextViewModel.Data))
		result, err := deliveryGRPC.ListText(ctx, listReqText)
		assertion.NoError(err)
		assertion.Equal(result.Data[0], listResText.Data[0])
		assertion.Equal(result.Total, totalText)
	})

	t.Run("negative list text", func(t *testing.T) {
		result, err := deliveryGRPC.ListText(context.Background(), listReqText)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})
}
