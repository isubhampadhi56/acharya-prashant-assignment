package jwtauth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret     []byte
	expiryTime int64
}

func (j *JWTManager) CreateToken(claims jwt.MapClaims) (string, error) {
	claims["exp"] = j.expiryTime
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "Auth-Server-1"
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := accessToken.SignedString(j.secret)
	if err != nil {
		return "", err
	}
	token = "Bearer " + token
	return token, nil
}
func (j *JWTManager) VerifyToken(jwtToken string) (jwt.MapClaims, error) {
	if jwtToken != "" && strings.HasPrefix(jwtToken, "Bearer ") {
		jwtToken = jwtToken[7:]
	} else {
		return nil, fmt.Errorf("invalid token or claims")
	}
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token or claims")
}

func InitializeJWTManager(secret []byte, expiryTime int64) JWT {
	return &JWTManager{
		secret:     secret,
		expiryTime: expiryTime,
	}
}
