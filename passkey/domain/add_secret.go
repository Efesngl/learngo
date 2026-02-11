package domain

import "encoding/base64"

type AddSecret struct {
	store     ISecretStore
	encrypter ICryptService
}

func NewAddSecret(store ISecretStore, encrypter ICryptService) *AddSecret {
	return &AddSecret{store: store, encrypter: encrypter}
}

func (as *AddSecret) Execute(name string, value string) error {
	if name == "" {
		return ErrEmptyName
	}
	if value == "" {
		return ErrEmptyValue
	}
	exists, err := as.store.Exists(name)
	if err != nil {
		return err
	}

	if exists {
		return ErrDuplicate
	}
	data := []byte(value)
	cipherText, err := as.encrypter.Encrypt(data)
	if err != nil {
		return err
	}
	secret := Secret{
		Name:  name,
		Value: base64.StdEncoding.EncodeToString(cipherText),
	}
	return as.store.Save(secret)
}
