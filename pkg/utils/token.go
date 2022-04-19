package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", t.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidToken(r *http.Request) (*jwt.Token, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}
	return token, nil
}

func UseToken(r *http.Request) (jwt.MapClaims, bool) {
	token, err := ValidToken(r)
	if err != nil {
		panic(err)
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(ok)
	}
	return claim, ok
}
