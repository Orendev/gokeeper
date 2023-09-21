package grpc

import (
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"google.golang.org/grpc"
)

type Options struct {
	Host string `env:"GRPC_HOST" env-default:"localhost"`
	Port uint   `env:"GRPC_PORT" env-default:"3200"`
}

type Repository struct {
	protobuff.KeeperServiceClient
	options Options
}

func New(
	cc *grpc.ClientConn,
	o Options,
) *Repository {
	service := protobuff.NewKeeperServiceClient(cc)

	d := &Repository{
		KeeperServiceClient: service,
	}

	d.SetOptions(o)

	return d
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}

//
//func (d *DeliveryClient) Run() (protobuff.KeeperServiceClient, error) {
//	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", d.options.Host, d.options.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		return nil, err
//	}
//
//	return protobuff.NewKeeperServiceClient(conn), nil
//}
//
//// ShutDown graceful stops the server.
//func (d *DeliveryClient) ShutDown() error {
//	d.serv.GracefulStop()
//	return nil
//}
