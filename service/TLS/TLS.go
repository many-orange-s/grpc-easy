package TLS

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-easy/config"
	"grpc-easy/interceptor/oauth"
	"io/ioutil"
	"log"
)

func CreateOp() []grpc.ServerOption {
	cert, err := tls.LoadX509KeyPair(config.Con.CarFile, config.Con.KeyFile)
	if err != nil {
		log.Fatalf("failed to load key part : %s", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(config.Con.CaFile)
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
		//grpc.UnaryInterceptor(basic.EnsureValidBasic),
		grpc.UnaryInterceptor(oauth.EnsureValid),
		grpc.StreamInterceptor(oauth.EnsureStreamAuth),
	}
	return opt
}
