package TLS

import (
	"client/config"
	"client/mytoken/oAuth"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"io/ioutil"
	"log"
)

func CreateOp() []grpc.DialOption {
	cert, err := tls.LoadX509KeyPair(config.Con.CrtFile, config.Con.KeyFile)
	if err != nil {
		log.Fatalf("Could not load client keu pair : %s", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(config.Con.CaFile)
	if err != nil {
		log.Fatalf("could not read ca cer %s", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("faild to append ca certs")
	}

	//auth := auth2.BasicAuth{Secret: config.Con.secret}
	auth := oauth.NewOauthAccess(oAuth.FetchToken())

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   config.Con.Hostname,
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
		})),
		grpc.WithPerRPCCredentials(auth),
	}
	return opts
}
