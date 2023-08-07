package helpers

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateSignedJWT(claims jwt.Claims) (string, error) {
	privateKeyPEM := strings.TrimSpace(os.Getenv("PRIVATE_KEY"))

	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.New("failed to decode PEM block")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	s, err := t.SignedString(privateKey)

	if err != nil {
		return "", err
	}

	return s, nil
}
