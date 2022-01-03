package oAuth

import (
	"client/config"
	"golang.org/x/oauth2"
)

func FetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: config.Con.Secrete,
	}
}
