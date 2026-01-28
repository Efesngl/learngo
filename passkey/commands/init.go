package commands

import (
	"github.com/Efesngl/learngo/passkey/storage"
)

func Init() error {
	saltStorage := storage.NewSaltStorage("salt.bin")

	exists, err := saltStorage.Exists()
	if err != nil {
		return nil
	}

	if exists {
		return nil
	}

	return saltStorage.Create()
}
