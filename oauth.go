package oauth

import (
	"net/http"

	"golang.org/x/oauth2"
)

// Oauth 授权接口
type Oauth interface {
	Authorize() string
	Callback(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Client(token *oauth2.Token) *http.Client
}
