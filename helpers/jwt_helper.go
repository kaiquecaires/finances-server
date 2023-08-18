package helpers

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	privateKeyPEM := strings.TrimSpace(os.Getenv("PRIVATE_KEY"))
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		// Check if the token is expired
		if expirationTime, ok := (*claims)["ExpiresAt"].(time.Time); ok {
			if time.Now().Unix() > expirationTime.Unix() {
				return nil, fmt.Errorf("token has expired")
			}
		}

		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

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
