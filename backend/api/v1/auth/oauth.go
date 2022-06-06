package auth

import (
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/utils"
	"encoding/json"
	"net/http"

	"io/ioutil"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/gitlab"
	"golang.org/x/oauth2/google"
)

type GithubEmailData struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo"

const oauthGithubUrlAPI = "https://api.github.com/user"

const oauthGithubEmailUrlAPI = "https://api.github.com/user/emails"

const oauthGitlabUrlAPI = "https://gitlab.example.com/api/v4/user"

const oauthGitlabEmailUrlAPI = "https://gitlab.example.com/api/v4/user/emails"

func createCallbackUrl(provider string) string {
	return "http://" + utils.GetEnv("HOST", "localhost") + ":" + utils.GetEnv("PORT", "9000") + "/api/v1/auths/" + provider + "/callback"
}

func GetAuthConfig(provider string) *oauth2.Config {
	switch provider {
	case "google":
		return &oauth2.Config{
			RedirectURL:  createCallbackUrl(provider),
			ClientID:     utils.GetEnv("GOOGLE_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GOOGLE_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
	case "github":
		return &oauth2.Config{
			RedirectURL:  createCallbackUrl(provider),
			ClientID:     utils.GetEnv("GH_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GH_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"user:email", "read:user"},
			Endpoint:     github.Endpoint,
		}
	case "gitlab":
		return &oauth2.Config{
			RedirectURL:  createCallbackUrl(provider),
			ClientID:     utils.GetEnv("GL_OAUTH_CLIENT_ID", ""),
			ClientSecret: utils.GetEnv("GL_OAUTH_CLIENT_SECRET", ""),
			Scopes:       []string{"read_user", "profile", "email"},
			Endpoint:     gitlab.Endpoint,
		}
	}
	return nil
}

func getUserDataClient(provider string, token *oauth2.Token) (*http.Request, *http.Request) {
	var req *http.Request
	var reqEmail *http.Request
	switch provider {
	case "google":
		req, _ = http.NewRequest("GET", oauthGoogleUrlAPI, nil)
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		return req, nil
	case "github":
		req, _ = http.NewRequest("GET", oauthGithubUrlAPI, nil)
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		reqEmail, _ = http.NewRequest("GET", oauthGithubEmailUrlAPI, nil)
		reqEmail.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		return req, reqEmail
	case "gitlab":
		req, _ = http.NewRequest("GET", oauthGitlabUrlAPI, nil)
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		reqEmail, _ = http.NewRequest("GET", oauthGitlabEmailUrlAPI, nil)
		reqEmail.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
		return req, reqEmail
	}
	return nil, nil
}

func GetUserData(provider string, token *oauth2.Token) (*models.UserData, error) {
	req, reqEmail := getUserDataClient(provider, token)
	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userData := new(models.UserData)
	if err := json.Unmarshal(contents, userData); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if reqEmail != nil {
		httpClient := &http.Client{}
		responseEmail, err := httpClient.Do(reqEmail)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer responseEmail.Body.Close()
		contents, err := ioutil.ReadAll(responseEmail.Body)
		emailData := []GithubEmailData{}
		if err := json.Unmarshal(contents, &emailData); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		userData.Email = emailData[0].Email
	}
	return userData, nil
}
