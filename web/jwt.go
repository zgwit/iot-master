package web

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

const JwtKey = "iot-master"
const JwtExpire = time.Hour * 24 * 30

func JwtGenerate(id string) (string, error) {
	var claims Claims
	claims.Id = id
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(JwtExpire))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func JwtVerify(str string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(str, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}
