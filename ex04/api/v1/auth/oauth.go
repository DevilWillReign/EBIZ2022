package auth

import (
	"apprit/store/api/v1/utils"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/slack"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo"

const oauthGithubUrlAPI = "https://api.github.com/user"

const oauthSlackUrlAPI = "https://api.github.com/user"

func createCallbackUrl(provider string) string {
	return "http://" + utils.GetEnv("HOST", "localhost") + ":" + utils.GetEnv("PORT", "9000") + "/api/v1/auths/" + provider + "/callback"
}

func GetAuthConfig(provider string) *oauth2.Config {
	switch provider {
	case "google":
		return &oauth2.Config{
			RedirectURL:  createCallbackUrl("google"),
			ClientID:     utils.GetEnv("GOOGLE_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GOOGLE_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	case "github":
		return &oauth2.Config{
			RedirectURL:  createCallbackUrl("github"),
			ClientID:     utils.GetEnv("GITHUB_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GITHUB_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		}
	case "slack":
		return &oauth2.Config{
			RedirectURL:  createCallbackUrl("slack"),
			ClientID:     utils.GetEnv("SLACK_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("SLACK_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"users:read.email"},
			Endpoint:     slack.Endpoint,
		}
	}
	return nil
}

func GetUserDataClient(provider string, token *oauth2.Token) *http.Request {
	var req *http.Request
	switch provider {
	case "google":
		req, _ = http.NewRequest("GET", oauthGoogleUrlAPI, nil)
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		return req
	case "github":
		req, _ = http.NewRequest("GET", oauthGithubUrlAPI, nil)
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		return req
	case "slack":
		req, _ = http.NewRequest("GET", oauthSlackUrlAPI, nil)
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		return req
	}
	return nil
}
