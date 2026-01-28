package crypt

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"

)

type ISaltProvider interface {
	Get() ([]byte, error)
}
type MasterKeyService struct {
	saltStore ISaltProvider
}

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
