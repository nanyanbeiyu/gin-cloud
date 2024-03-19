package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"go_cloud/util/conf"
	"time"
)

type MyCustomClaims struct {
	UserName string
	jwt.RegisteredClaims
}

func SetJWTToken(userName string) string {
	claims := MyCustomClaims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-cloud",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		panic(err)
	}
	return signedString
}
