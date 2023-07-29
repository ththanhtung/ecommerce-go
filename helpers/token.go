package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	UserID string
	Email string
}

// create private, public token pair rsa
func GenerateKeyPair() (privateKeyPEM, publicKeyPEM []byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	if err = privateKey.Validate(); err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	// Convert private key to PEM
	privateKeyPEM = pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	// Convert public key to PEM
	publicKeyPEM = pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	)

	return privateKeyPEM, publicKeyPEM, nil
}

type JWT struct {
	privateKey []byte
	publicKey  []byte
}

func NewJWT(privateKey, publicKey []byte) *JWT {
	return &JWT{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (j JWT) Sign(content any, ttl time.Duration) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["dat"] = content
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j JWT) Validate(token string) (any, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return nil, err
	}

	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("validate: invalid")
	}

	return claims["dat"], nil
}

func CreateTokensPair(payload any, privateKey, publicKey []byte) (string, string, error) {
	jwt := NewJWT(privateKey, publicKey)
	accessToken, err := jwt.Sign(payload, 2*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.Sign(payload, 5*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
