package repository

import "github.com/pkg/errors"

var (
	ErrIncorrectPassword = errors.New("incorrect email/password")
	ErrDataNotFound      = errors.Errorf("data not found")
	ErrCardNotFound      = errors.Errorf("card not found")
	ErrAccountNotFound   = errors.Errorf("account not found")
	ErrNotFoundUser      = errors.New("user not found")
	ErrNoAuth            = errors.New("grpc: authorization has not been completed yet, enter your username and password")
)
