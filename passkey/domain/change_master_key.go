package domain

import (
	"encoding/base64"
)

type ChangeMasterKey struct {
	store           ISecretStore
	oldCryptService ICryptService
	newCryptService ICryptService
}

func NewChangeMasterKey(store ISecretStore, oldCryptService ICryptService, newCryptService ICryptService) *ChangeMasterKey {
	return &ChangeMasterKey{store: store, oldCryptService: oldCryptService, newCryptService: newCryptService}
}

func (c *ChangeMasterKey) Execute() error {
	secrets, err := c.store.List()
	if err != nil {
		return err
	}
	for index, secret := range secrets {
		oldValue, err := base64.StdEncoding.DecodeString(secret.Value)
		if err != nil {
			return err
		}
		decryptedValue, err := c.oldCryptService.Decrypt(oldValue)
		if err != nil {
			return err
		}
		newEncryptedvalue, err := c.newCryptService.Encrypt(decryptedValue)
		if err != nil {
			return err
		}
		secrets[index].Value = base64.StdEncoding.EncodeToString(newEncryptedvalue)
	}
	return c.store.SaveAll(secrets)
}
