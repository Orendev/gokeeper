package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/Orendev/gokeeper/internal/app/client/configs"
	"github.com/Orendev/gokeeper/internal/app/client/repository/client/grpc/interceptors"
	"github.com/Orendev/gokeeper/pkg/protobuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	protobuff.KeeperServiceClient
	authInterceptor *interceptors.AuthInterceptor
	options         configs.ServerGRPC
}

func New(
	authInterceptor *interceptors.AuthInterceptor,
	o configs.ServerGRPC,
) (*Client, error) {
	var err error
	client := &Client{
		authInterceptor: authInterceptor,
	}

	client.SetOptions(o)

	err = client.Connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) SetOptions(options configs.ServerGRPC) {
	if c.options != options {
		c.options = options
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	//Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failedtoaddserverCA'scertificate")
	}

	//Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	//Create the credential sand return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func (c *Client) Connect() error {

	serverAddress := fmt.Sprintf("%s:%d", c.options.Host, c.options.Port)
	transportOption := grpc.WithInsecure()

	if c.options.EnabledTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return err
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	conn, err := grpc.Dial(
		serverAddress,
		transportOption,
		grpc.WithUnaryInterceptor(c.authInterceptor.UnaryAuth()),
	)

	if err != nil {
		return err
	}
	c.KeeperServiceClient = protobuff.NewKeeperServiceClient(conn)

	return nil
}
