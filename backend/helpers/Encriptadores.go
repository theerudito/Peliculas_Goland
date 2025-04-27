package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func EncryptString(value string) (string, error) {

	key := []byte("thisis32bitlongpassphraseimusing")

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

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(value), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptString(value string) (string, error) {

	key := []byte("thisis32bitlongpassphraseimusing")

	ciphertext, err := base64.StdEncoding.DecodeString(value)
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

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func ComparePassword(inputPassword, encryptedPassword string) (bool, error) {
	decryptedPassword, err := DecryptString(encryptedPassword)
	if err != nil {
		return false, err
	}

	return inputPassword == decryptedPassword, nil
}
