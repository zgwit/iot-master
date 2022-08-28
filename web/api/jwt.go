package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

const key = "iot-master"

func generate() (string, error) {
	var claims Claims
	claims.ExpiresAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func verify(str string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(str, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}

func jwtMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	token = c.Request.URL.Query().Get("token")
	claims, err := verify(token)
	if err != nil {
		c.Abort()
	}
	c.Set("claims", claims)
	c.Next()
}
