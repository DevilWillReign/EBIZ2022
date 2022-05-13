package auth

import (
	"apprit/store/api/v1/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/slack"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

const oauthGithubUrlAPI = "https://api.github.com/user"

const oauthSlackUrlAPI = "https://api.github.com/user"

func GetAuthConfig(provider string) (*oauth2.Config, string) {
	switch provider {
	case "google":
		return &oauth2.Config{
			RedirectURL:  "http://" + utils.GetEnv("HOST", "localhost") + ":" + utils.GetEnv("PORT", "9000") + "/api/v1/auths/google/callback",
			ClientID:     utils.GetEnv("GOOGLE_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GOOGLE_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}, oauthGoogleUrlAPI
	case "github":
		return &oauth2.Config{
			RedirectURL:  "http://" + utils.GetEnv("HOST", "localhost") + ":" + utils.GetEnv("PORT", "9000") + "/api/v1//auths/github/callback",
			ClientID:     utils.GetEnv("GITHUB_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GITHUB_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		}, oauthGithubUrlAPI
	case "slack":
		return &oauth2.Config{
			RedirectURL:  "http://" + utils.GetEnv("HOST", "localhost") + ":" + utils.GetEnv("PORT", "9000") + "/api/v1/auths/slack/callback",
			ClientID:     utils.GetEnv("SLACK_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("SLACK_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"users:read.email"},
			Endpoint:     slack.Endpoint,
		}, oauthSlackUrlAPI
	}
	return nil, ""
}
