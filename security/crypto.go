package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func PrivateKeyFromString(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))

	if block == nil {
		return nil, errors.New("invalid private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func PrivateKeyToString(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return string(privateKeyPEM), nil
}

func Encrypt(publicKey *rsa.PublicKey, text string) (string, error) {
	cipherBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(text),
		nil,
	)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

func Decrypt(privateKey *rsa.PrivateKey, cipherText string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)

	if err != nil {
		return "", err
	}

	decryptedBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		cipherBytes,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(decryptedBytes), nil
}
