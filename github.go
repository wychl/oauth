package oauth

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/xid"
	"golang.org/x/oauth2"
)

// Github oauth impl
type Github struct {
	conf *oauth2.Config
}

var _ Oauth = &Github{}

// New Oauth instance
func New(conf *oauth2.Config) *Github {
	return &Github{
		conf: conf,
	}
}

// Authorize Oauth Impl
func (o *Github) Authorize() string {
	state := xid.New().String()
	return o.conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// Callback Oauth Impl
func (o *Github) Callback(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	return o.conf.Exchange(ctx, code, opts...)
}

// Client Oauth Impl
func (o *Github) Client(token *oauth2.Token) *http.Client {
	return o.conf.Client(context.Background(), token)
}
