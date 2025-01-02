package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OauthConfig *oauth2.Config

func InitOAuth() {
	// Define OAuth2 config, with credentials from Google Cloud Console
	OauthConfig = &oauth2.Config{
		ClientID:     "Your_Client_ID",
		ClientSecret: "Your_Client_Secret",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/gmail.readonly",
			"https://www.googleapis.com/auth/drive.readonly",
		},
		RedirectURL: "http://localhost:8888/api/v1/admin/callback",
		Endpoint:    google.Endpoint,
	}
}
