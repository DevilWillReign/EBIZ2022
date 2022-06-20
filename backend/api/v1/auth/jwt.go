package auth

import (
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	models.UserData
	jwt.StandardClaims
}

func GetCustomClaimsConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:      &JwtCustomClaims{},
		SigningKey:  []byte(utils.GetEnv("JWT_SECRET", "secret")),
		TokenLookup: "header:Authorization,cookie:userinfo",
	}
}

func CreateToken(userData models.UserData) (string, error) {
	claims := &JwtCustomClaims{
		models.UserData{
			ID:    userData.ID,
			Name:  userData.Name,
			Email: userData.Email,
			Admin: userData.Admin,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(utils.GetEnv("JWT_SECRET", "secret")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func IsNotAdmin(c echo.Context) bool {
	userToken := c.Get("user").(*jwt.Token)
	if userToken == nil {
		return true
	}
	claims := userToken.Claims.(*JwtCustomClaims)
	return !claims.Admin
}

func IsNotAdminAndSameUser(c echo.Context, id uint64) bool {
	userToken := c.Get("user").(*jwt.Token)
	if userToken == nil {
		return true
	}
	claims := userToken.Claims.(*JwtCustomClaims)
	return !claims.Admin && claims.ID != uint(id)
}

func GetUserIdOrError(c echo.Context) (uint64, error) {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return 0, echo.ErrUnauthorized
	}
	claims := user.Claims.(*JwtCustomClaims)
	return uint64(claims.ID), nil
}
