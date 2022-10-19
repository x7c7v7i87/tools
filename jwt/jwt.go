package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("!@)*#)!@U#@*!@!)!!!!")

type Claims struct {
	Email    string `json:"email"`
	Password string `json:"passwd"`
	jwt.StandardClaims
}

func UpdateTokenTime() int64 {
	nowTime := time.Now()
	//hour:=60 * 60 * 3
	expireTime := nowTime.Add(3 * time.Hour)
	newTimeInt := expireTime.Unix()
	return newTimeInt
}

func GenerateToken(email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tools",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
