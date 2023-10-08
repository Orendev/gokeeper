package binary

import (
	"context"
	"os"
	"testing"

	"github.com/Orendev/gokeeper/internal/pkg/domain/binary"
	"github.com/Orendev/gokeeper/internal/pkg/repository"
	"github.com/Orendev/gokeeper/internal/pkg/repository/storage/mocks"
	"github.com/Orendev/gokeeper/pkg/type/queryParameter"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	storageRepository   = new(mocks.Storage)
	ucBinary            *UseCase
	data                = make(map[uuid.UUID]*binary.BinaryData)
	binaries            []*binary.BinaryData
	listBinaryViewModel = &binary.ListBinaryViewModel{}
	createBinary        *binary.BinaryData
	updateBinary        *binary.BinaryData
	parameter           = queryParameter.QueryParameter{}
	total               uint64
)

func TestMain(m *testing.M) {
	titleObj, _ := title.New("Test title")
	userID := uuid.New()

	createBinary, _ = binary.New(
		userID,
		*titleObj,
		[]byte("binary data"),
		[]byte("test comment"),
	)

	titleUpdateObj, _ := title.New("Test Update")

	updateBinary, _ = binary.New(
		userID,
		*titleUpdateObj,
		[]byte("binary data"),
		[]byte("test comment"),
	)

	parameter.Pagination.Limit = 1
	parameter.Pagination.Offset = 0
	total = 1

	os.Exit(m.Run())
}

func initTestUseCaseBinary(t *testing.T) {
	assertion := assert.New(t)

	storageRepository.On("CreateBinary",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, binary *binary.BinaryData) *binary.BinaryData {
			assertion.Equal(binary, createBinary)
			binaries = append(binaries, binary)
			listBinaryViewModel.Data = binaries
			data[binary.ID()] = binary

			return binary
		}, func(ctx context.Context, binary *binary.BinaryData) error {
			return nil
		})

	storageRepository.On("UpdateBinary",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, binary *binary.BinaryData) *binary.BinaryData {
			assertion.Equal(binary, updateBinary)
			data[binary.ID()] = binary
			listBinaryViewModel.Data = binaries

			return binary
		}, func(ctx context.Context, binary *binary.BinaryData) error {
			return nil
		})

	storageRepository.On("DeleteBinary",
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

	storageRepository.On("ListBinary",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryParameter.QueryParameter) *binary.ListBinaryViewModel {
			t := len(listBinaryViewModel.Data)
			listBinaryViewModel.Limit = parameter.Pagination.Limit
			listBinaryViewModel.Offset = parameter.Pagination.Offset
			listBinaryViewModel.Total = uint64(t)

			return listBinaryViewModel

		}, func(ctx context.Context, parameter queryParameter.QueryParameter) error {
			return nil
		})
}

func TestBinary(t *testing.T) {
	initTestUseCaseBinary(t)
	ucBinary = New(storageRepository, Options{})
	assertion := assert.New(t)

	t.Run("create binary", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucBinary.Create(ctx, createBinary)
		assertion.NoError(err)
		assertion.Equal(result, createBinary)
	})

	t.Run("update binary", func(t *testing.T) {
		var ctx = context.Background()

		result, err := ucBinary.Update(ctx, updateBinary)
		assertion.NoError(err)
		assertion.Equal(result, updateBinary)
	})

	t.Run("delete binary", func(t *testing.T) {
		var ctx = context.Background()

		err := ucBinary.Delete(ctx, createBinary.ID())
		assertion.NoError(err)

	})

	t.Run("list binary", func(t *testing.T) {
		var ctx = context.Background()
		total = uint64(len(listBinaryViewModel.Data))
		result, err := ucBinary.List(ctx, parameter)
		assertion.NoError(err)
		assertion.Equal(result.Data, listBinaryViewModel.Data)
		assertion.Equal(result.Total, total)
	})
}
