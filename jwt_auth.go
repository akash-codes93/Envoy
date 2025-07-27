package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJwtToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":       time.Now().Unix(),
		"uid":       user.ID,
		"device_id": "test-device-id",
		"expiry":    time.Now().Add(time.Minute * 5).Unix(),
		"role":      "Listner",
		"version":   "v2",
		"category":  "access",
		"tenant":    "pocket_fm",
		"platform":  "android",
	})

	return token.SignedString([]byte("iam_auth_secret"))
}
