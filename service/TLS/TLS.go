package TLS

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-easy/interceptor/basic"
	"io/ioutil"
	"log"
)

const (
	crtFile = "D:\\go_project\\keys\\server.crt"
	keyFile = "D:\\go_project\\keys\\server.key"
	caFile  = "D:\\go_project\\keys\\ca.crt"
)

func CreateOp() []grpc.ServerOption {
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key part : %s", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate :%s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("faild to append ca certificate")
	}

	opt := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAnyClientCert,
			Certificates: []tls.Certificate{cert},
			ClientCAs:    certPool,
		})),
		grpc.UnaryInterceptor(basic.EnsureValidBasic),
	}
	return opt
}
