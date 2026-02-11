package domain

import (
	"encoding/base64"
	"errors"
)

type GetSecret struct {
	store     ISecretStore
	encrypter ICryptService
}

func NewGetSecret(store ISecretStore, encrypter ICryptService) *GetSecret {
	return &GetSecret{store: store, encrypter: encrypter}
}

func (gs *GetSecret) Execute(name string) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	_, err := gs.store.Exists(name)
	if err != nil {
		return "", err
	}
	secret, err := gs.store.Get(name)
	if err != nil {
		return "", err
	}
	cipheredText, err := base64.StdEncoding.DecodeString(secret.Value)
	if err != nil {
		return "", err
	}
	data, err := gs.encrypter.Decrypt(cipheredText)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var (
	ErrSecretIsNotExists = errors.New("Secret is not exists")
)
