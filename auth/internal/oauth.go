package internal

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     "GOOGLE_CLIENT_ID",
		ClientSecret: "GOOGLE_CLIENT_SECRET",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

func GetGoogleOAuthURL(state string) string {
	return GoogleOauthConfig.AuthCodeURL(state)
}

func ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return GoogleOauthConfig.Exchange(ctx, code)
}
