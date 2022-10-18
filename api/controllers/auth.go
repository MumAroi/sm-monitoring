package controllers

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     "35daa01587ea64bdea10",
		ClientSecret: "b1fd92f3b2e04f5e8dbb9bd368449582ce65b85f",
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")
}
