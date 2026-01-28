package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

type AES struct {
	Masterkey []byte
}

func NewAESEncrypter(masterkey []byte) *AES {
	return &AES{
		Masterkey: masterkey,
	}
}

func (a *AES) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Masterkey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nil, nonce, data, nil)
	return append(nonce, cipherText...), nil
}

func (a *AES) Decrypt(cipherData []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Masterkey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(cipherData) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce := cipherData[:nonceSize]
	encryptedText := cipherData[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, encryptedText, nil)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}
