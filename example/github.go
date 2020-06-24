package main

import (
	"net/http"
	"os"

	"github.com/wychl/oauth"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var conf = &oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Endpoint:     github.Endpoint,
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	Scopes:       []string{},
}

var oathClient oauth.Oauth = oauth.New(conf)

func main() {
	router := gin.Default()
	router.GET("/auth", auth)
	router.GET("/token", token)
	http.ListenAndServe(":9999", router)
}

func auth(c *gin.Context) {
	uri := oathClient.Authorize()
	c.Redirect(http.StatusSeeOther, uri)
}

func token(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	var (
		token *oauth2.Token
		err   error
	)

	token, err = oathClient.Callback(code, oauth2.SetAuthURLParam("state", state))
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code": 1,
			"resp": err,
		})
		return
	}

	c.JSON(http.StatusOK, token)
}

// authorize http://localhost:9999/auth
