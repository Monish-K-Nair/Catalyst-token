package utils

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID            string
	Email         string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func CreateJWT(Data map[string]interface{}, secret string) (string, error) {
	expiredAt := time.Now().AddDate(0, 0, 7).Unix()
	claims := jwt.MapClaims{}
	claims["authorization"] = true
	claims["exp"] = expiredAt
	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte("12345678"))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyJWTToken(c *gin.Context, secret string) (*jwt.Token, error) {
	val := strings.Trim(strings.SplitAfter(c.GetHeader("Authorization"), "Bearer")[1], " ")
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return []byte("12345678"), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeJWTToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
