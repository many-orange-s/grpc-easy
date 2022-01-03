package TLS

import (
	"client/mytoken/oAuth"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"io/ioutil"
	"log"
)

const (
	hostname = "many.orange"
	crtFile  = "D:\\go_project\\keys\\client.crt"
	keyFile  = "D:\\go_project\\keys\\client.key"
	caFile   = "D:\\go_project\\keys\\ca.crt"
	Secrete  = "太阳高高我要起早"
)

func CreateOp() []grpc.DialOption {
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("Could not load client keu pair : %s", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca cer %s", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("faild to append ca certs")
	}

	//auth := auth2.BasicAuth{Secret: Secrete}
	auth := oauth.NewOauthAccess(oAuth.FetchToken())

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostname,
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
		})),
		grpc.WithPerRPCCredentials(auth),
	}
	return opts
}
