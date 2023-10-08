package grpc

import (
	"fmt"
	"net"

	"github.com/Orendev/gokeeper/internal/app/server/delivery/grpc/interceptors"
	"github.com/Orendev/gokeeper/internal/pkg/useCase"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"google.golang.org/grpc"
)

type Options struct {
	Host string `env:"GRPC_HOST" env-default:"localhost"`
	Port uint   `env:"GRPC_PORT" env-default:"3200"`
}

type Delivery struct {
	protobuff.UnimplementedKeeperServiceServer
	ucUser    useCase.User
	ucAccount useCase.Account
	ucCard    useCase.Card
	ucText    useCase.Text
	ucBinary  useCase.Binary
	serv      *grpc.Server

	jwtManager *auth.JWTManager
	options    Options
}

func New(
	ucUser useCase.User,
	ucAccount useCase.Account,
	ucCard useCase.Card,
	ucText useCase.Text,
	ucBinary useCase.Binary,
	jwtManager *auth.JWTManager,
	o Options) *Delivery {

	var d = &Delivery{
		ucUser:     ucUser,
		ucAccount:  ucAccount,
		ucCard:     ucCard,
		ucText:     ucText,
		ucBinary:   ucBinary,
		jwtManager: jwtManager,
	}

	d.SetOptions(o)

	interceptor := interceptors.NewAuthInterceptor(jwtManager, auth.AccessibleRoles())

	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.ChainUnaryInterceptor(interceptor.UnaryLogger()),
		grpc.ChainUnaryInterceptor(interceptor.UnaryAuth()),
	)

	s := grpc.NewServer(
		opts...,
	)

	protobuff.RegisterKeeperServiceServer(s, d)

	d.serv = s

	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", d.options.Host, d.options.Port))
	if err != nil {
		return err
	}
	return d.serv.Serve(listen)
}

// ShutDown graceful stops the server.
func (d *Delivery) ShutDown() {
	d.serv.GracefulStop()
}
