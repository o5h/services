package token

import (
	"crypto/rand"
	"encoding/base64"
)

var (
	JWT_KEY           = []byte("your-secret-key")
	SUBJECT_TOKEN_KEY = []byte("your-secret-key")
)

// GenerateRandomKey generates random secret key for JWT.
// For JWT, a key length of 256 bits (32 bytes) or more is commonly recommended.
// Adjust the keyLength variable according to your security requirements.
// Additionally, ensure that you store the secret key securely and do not expose it in your code or share it publicly.
func GenerateRandomKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomKey := base64.URLEncoding.EncodeToString(randomBytes)
	return randomKey, nil
}
