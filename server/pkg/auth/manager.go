package auth

import (
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Manager struct {
	tokenKey      string
	signingKey    []byte
	signingMethod jwt.SigningMethod
}

var manager *Manager

func InitManager(tokenKey, signingKey string) *Manager {
	rand.Seed(time.Now().UnixNano())
	manager = &Manager{
		tokenKey:      tokenKey,
		signingKey:    []byte(signingKey),
		signingMethod: jwt.SigningMethodHS256,
	}
	return manager
}

func GetManager() *Manager {
	return manager
}