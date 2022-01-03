package oAuth

import (
	"golang.org/x/oauth2"
)

func FetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "太阳高高我要起早",
	}
}
