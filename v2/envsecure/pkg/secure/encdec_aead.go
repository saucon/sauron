package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/pkg/errors"
	"io"
	"os"
)

// encryptAEAD encrypts text using AES-GCM with the provided secret
func encryptAEAD(text, MySecret string) (string, error) {
	key := []byte(MySecret)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16, 24, or 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, []byte(text), nil)
	full := append(nonce, cipherText...)
	return base64.StdEncoding.EncodeToString(full), nil
}

// decryptAEAD decrypts base64-encoded AES-GCM encrypted text with the provided secret
func decryptAEAD(encryptedText, MySecret string) (string, error) {
	key := []byte(MySecret)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16, 24, or 32 bytes")
	}

	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < aesGCM.NonceSize() {
		return "", errors.New("invalid encrypted data")
	}

	nonce := data[:aesGCM.NonceSize()]
	cipherText := data[aesGCM.NonceSize():]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// encryptRSA encrypts text using an RSA public key (in PEM format)
func encryptRSA(text, publicKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "PUBLIC KEY" {
		return "", errors.New("invalid public key PEM")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("not RSA public key")
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(text))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// decryptRSA decrypts base64-encoded RSA encrypted text using a private key (in PEM format)
func decryptRSA(encryptedText, privateKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("invalid private key PEM")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPriv, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not RSA private key")
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPriv, cipherBytes)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func ReadKeyFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadKeyFromBytes(key []byte) (string, error) {
	return string(key), nil
}
