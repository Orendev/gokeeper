package grpc

import (
	"testing"
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	useCaseBinary "github.com/Orendev/gokeeper/internal/pkg/useCase/binary"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

var (
	// binaries
	binaries            []*binary.BinaryData
	dataBinary          = make(map[uuid.UUID]*binary.BinaryData)
	listBinaryViewModel = &binary.ListBinaryViewModel{}
	createBinary        *binary.BinaryData
	updateBinary        *binary.BinaryData
	totalBinary         uint64
	createReqBin        *protobuff.CreateBinaryRequest
	createResBin        *protobuff.CreateBinaryResponse

	updateReqBin *protobuff.UpdateBinaryRequest
	updateResBin *protobuff.UpdateBinaryResponse

	deleteReqBin *protobuff.DeleteBinaryRequest
	deleteResBin *protobuff.DeleteBinaryResponse

	listReqBin *protobuff.ListBinaryRequest
	//listResBin *protobuff.ListBinaryResponse
)

func testMainBinary() {
	userID := uuid.New()
	//accessToken, _ := jwtManager.Generate(userID)
	//header := metadata.New(map[string]string{auth.AuthorizationKey: accessToken})
	//ctx = metadata.NewIncomingContext(context.Background(), header)

	titleObj, _ := title.New("Test title")

	createBinary, _ = binary.New(
		userID,
		*titleObj,
		[]byte("test ya.ru"),
		[]byte("test comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0

	binData := &protobuff.Data{
		ID:        createBinary.ID().String(),
		Title:     createBinary.Title().String(),
		Data:      createBinary.Data(),
		Comment:   createBinary.Comment(),
		CreatedAt: createBinary.CreatedAt().Format(time.RFC3339),
		UpdatedAt: createBinary.UpdatedAt().Format(time.RFC3339),
		UserID:    createBinary.UserID().String(),
	}
	createReqBin = &protobuff.CreateBinaryRequest{Data: binData}

	createResBin = &protobuff.CreateBinaryResponse{
		ID: createBinary.ID().String(),
	}

	titleUpdateObj, _ := title.New("Test Update")

	updateBinary, _ = binary.NewWithID(
		createBinary.ID(),
		userID,
		*titleUpdateObj,
		[]byte("text data update"),
		[]byte("test comment update"),
		createBinary.CreatedAt().UTC(),
		time.Now().UTC(),
	)

	binUpdateData := &protobuff.Data{
		ID:        updateBinary.ID().String(),
		Title:     updateBinary.Title().String(),
		Data:      updateBinary.Data(),
		Comment:   updateBinary.Comment(),
		CreatedAt: updateBinary.CreatedAt().Format(time.RFC3339),
		UpdatedAt: updateBinary.UpdatedAt().Format(time.RFC3339),
		UserID:    updateBinary.UserID().String(),
	}

	updateReqBin = &protobuff.UpdateBinaryRequest{Data: binUpdateData}

	updateResBin = &protobuff.UpdateBinaryResponse{
		ID: updateBinary.ID().String(),
	}

	deleteReqBin = &protobuff.DeleteBinaryRequest{
		ID: updateBinary.ID().String(),
	}
	deleteResBin = &protobuff.DeleteBinaryResponse{
		ID: updateBinary.ID().String(),
	}

	listReqBin = &protobuff.ListBinaryRequest{
		Limit:  parameter.Pagination.Limit,
		Offset: parameter.Pagination.Offset,
	}

	totalBinary = 1

	binListData := []*protobuff.Data{}

	binListData = append(binListData, binUpdateData)
	//listResBin = &protobuff.ListBinaryResponse{
	//	Limit:  parameter.Pagination.Limit,
	//	Offset: parameter.Pagination.Offset,
	//	Total:  totalBinary,
	//	Data:   binListData,
	//}
}

func initTestUseCaseBinary(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateBinary",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, binary *binary.BinaryData) *binary.BinaryData {
			assertion.Equal(binary.ID(), updateBinary.ID())
			binaries = append(binaries, binary)
			listBinaryViewModel.Data = binaries
			dataBinary[binary.ID()] = binary

			return binary
		}, func(ctx context.Context, binary *binary.BinaryData) error {
			return nil
		})

	storageRepository.On("UpdateBinary",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, binary *binary.BinaryData) *binary.BinaryData {
			assertion.Equal(binary.ID(), updateBinary.ID())
			if len(binaries) > 0 {
				binaries[0] = binary
			} else {
				binaries = append(binaries, binary)
			}

			dataBinary[binary.ID()] = binary
			listBinaryViewModel.Data = binaries

			return binary
		}, func(ctx context.Context, binary *binary.BinaryData) error {
			return nil
		})

	storageRepository.On("DeleteBinary",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, id uuid.UUID) error {
			if _, ok := dataBinary[id]; !ok {
				return repository.ErrDataNotFound
			}
			return nil
		}, func(ctx context.Context, id uuid.UUID) error {
			return nil
		})

	storageRepository.On("ListBinary",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) *binary.ListBinaryViewModel {

			listBinaryViewModel.Limit = parameter.Pagination.Limit
			listBinaryViewModel.Offset = parameter.Pagination.Offset
			listBinaryViewModel.Total = uint64(len(listBinaryViewModel.Data))

			return listBinaryViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestDeliveryBinary(t *testing.T) {
	initTestUseCaseBinary(t)
	ucBinary = useCaseBinary.New(storageRepository, useCaseBinary.Options{})
	option := Options{}
	deliveryGRPC := New(ucUser, ucAccount, ucCard, ucText, ucBinary, jwtManager, option)
	assertion := assert.New(t)

	t.Run("positive create binary", func(t *testing.T) {
		result, err := deliveryGRPC.CreateBinary(ctx, createReqBin)
		assertion.NoError(err)
		assertion.Equal(result, createResBin)
	})

	t.Run("negative create binary", func(t *testing.T) {
		result, err := deliveryGRPC.CreateBinary(context.Background(), createReqBin)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive update binary", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateBinary(ctx, updateReqBin)
		assertion.NoError(err)
		assertion.Equal(result, updateResBin)
	})

	t.Run("negative update binary", func(t *testing.T) {
		result, err := deliveryGRPC.UpdateBinary(context.Background(), updateReqBin)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})

	t.Run("positive delete binary", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteBinary(ctx, deleteReqBin)
		assertion.NoError(err)
		assertion.Equal(result, deleteResBin)
	})

	t.Run("negative delete binary", func(t *testing.T) {
		result, err := deliveryGRPC.DeleteBinary(context.Background(), &protobuff.DeleteBinaryRequest{ID: uuid.New().String()})
		assertion.Errorf(err, "")
		assertion.Nil(result)
	})

	t.Run("positive list binary", func(t *testing.T) {
		totalBinary = uint64(len(listBinaryViewModel.Data))
		result, err := deliveryGRPC.ListBinary(ctx, listReqBin)
		assertion.NoError(err)
		assertion.Equal(result.Total, totalBinary)
	})

	t.Run("negative list binary", func(t *testing.T) {
		result, err := deliveryGRPC.ListBinary(context.Background(), listReqBin)
		assertion.Errorf(err, auth.ErrorTokenContextMissing.Error())
		assertion.Nil(result)
	})
}
