package helpers

import (
	"errors"
	"net/http"
	"server/models"
	"server/server"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GetTokenFromRequest(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		return nil, errors.New("no token found in request")
	}

	return token, nil
}