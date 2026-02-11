package commands

import (
	"fmt"

	"github.com/Efesngl/learngo/passkey/crypt"
	"github.com/Efesngl/learngo/passkey/domain"
	"github.com/Efesngl/learngo/passkey/storage"
)

func initStorages(slatStorage string, jsonStorage string) (*storage.SaltStorage, *storage.JSONStorage) {
	saltStorage := storage.NewSaltStorage(slatStorage)
	JsonStorage := storage.NewJSONStorage(jsonStorage)
	return saltStorage, JsonStorage
}

func deriveMasterKey(SaltStore domain.ISaltStore, message string) ([]byte, error) {
	fmt.Print(message)
	var password []byte
	fmt.Scan(&password)
	MasterKeyService := crypt.NewMasterKeyService(SaltStore)
	return MasterKeyService.Derive(password)
}

func getSecretValue(args []string) string {
	if len(args) >= 2 {
		return args[1]
	}

	fmt.Print("Please enter secret value: ")
	var secretValue string
	fmt.Scanln(&secretValue)
	return secretValue
}
