package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hongjie104/NAS-server/config"
)

var jwtSecret = []byte(config.Config.APP.JwtSecret)

// Claims a
type Claims struct {
	UserName string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

// GenerateToken a
func GenerateToken(username, id string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour * 365)
	claims := Claims{
		username,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hj",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken a
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
