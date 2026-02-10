package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

type ISaltProvider interface {
	Get() ([]byte, error)
}
type ISecretStore interface {
	First() (string, error)
}

type MasterKeyService struct {
	saltStore ISaltProvider
}

var ErrorInvalidMasterKey = errors.New("Invalid Master key")

func NewMasterKeyService(saltStore ISaltProvider) *MasterKeyService {
	return &MasterKeyService{
		saltStore: saltStore,
	}
}

func (mks *MasterKeyService) Derive(password []byte) ([]byte, error) {
	salt, err := mks.saltStore.Get()
	if err != nil {
		return nil, err
	}

	masterKey := pbkdf2.Key(password, salt, 4096, 32, sha256.New)
	return masterKey, nil
}

func (mks *MasterKeyService) Verify(masterKey []byte, store ISecretStore) error {
	secretValue, err := store.First()
	if err != nil {
		return err
	}

	cipherText, err := base64.StdEncoding.DecodeString(secretValue)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return errors.New("ciphertext too short")
	}

	nonce := cipherText[:nonceSize]
	encryptedText := cipherText[nonceSize:]

	_, err = gcm.Open(nil, nonce, encryptedText, nil)

	if err != nil {
		return ErrorInvalidMasterKey
	}
	return nil
}
