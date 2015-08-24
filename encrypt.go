package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Key used to encrypt/decrypt password
// Note: this is by no mean meant to be a bullet proof encryption scheme but just rather a way to
// avoid having sensitive data written in plain text on disk.
func seekret() []byte {
	return []byte(`xLtwITJ$tIh@*B#lQL53W!Qq`)
}

func encodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func decodeBase64(b []byte) []byte {
	data, _ := base64.StdEncoding.DecodeString(string(b))
	return data
}

// Encrypt encrypts the given text with a hard-coded secret. Not truly secure.
func Encrypt(text string) (string, error) {
	bytes := []byte(text)
	key := seekret()
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := encodeBase64(bytes)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return string(encodeBase64(ciphertext)), nil
}

// Decrypt decrypts the given encrypted string using the hard-coded secret.
func Decrypt(text string) (string, error) {
	if text == "" {
		return "", nil
	}
	key := seekret()
	bytes := decodeBase64([]byte(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(bytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := bytes[:aes.BlockSize]
	bytes = bytes[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(bytes, bytes)
	return string(decodeBase64(bytes)), nil
}
