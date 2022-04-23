package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func (m *Manager) CreateToken(val string) (string, error) {
	token := jwt.NewWithClaims(m.signingMethod, jwt.MapClaims{
		m.tokenKey: val,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(m.signingKey)
}

func (m *Manager) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.signingKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("error get claims from token")
	}
	return claims[m.tokenKey].(string), nil
}

func (m *Manager) ParseTokenId(tokenStr string) (int, error) {
	tokenVal, err := m.ParseToken(tokenStr)
	if err != nil {
		return 0, err
	}

	uid, err := strconv.ParseUint(tokenVal, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(uid), nil
}

func (m *Manager) TokenValid(r *http.Request) error {
	tokenStr := ExtractToken(r)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.signingKey, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	tokenList := strings.Split(bearerToken, "Bearer ")
	if len(tokenList) == 2 {
		return tokenList[1]
	}
	return ""
}

func (m *Manager) ExtractTokenID(r *http.Request) (int, error) {
	tokenStr := ExtractToken(r)
	return m.ParseTokenId(tokenStr)
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
