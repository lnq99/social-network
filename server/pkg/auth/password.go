package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (m *Manager) GetHashSalt(password string) (salt string, hashed string) {
	salt = randStringBytes(8)
	h := sha1.New()
	h.Write([]byte(salt + password))
	hashed = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}

func (m *Manager) ComparePassword(password, salt, hashed string) bool {
	h := sha1.New()
	h.Write([]byte(salt + password))
	res := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return res == hashed
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}