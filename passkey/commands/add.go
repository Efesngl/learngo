package commands

import (
	"fmt"

	"github.com/Efesngl/learngo/passkey/crypt"
	"github.com/Efesngl/learngo/passkey/domain"
)

func Add(args []string) {
	if len(args) < 1 {
		fmt.Println(domain.ErrEmptyName)
		return
	}

	saltStorage, JsonStorage := initStorages("salt.bin", "secrets.json")
	masterKeyService := crypt.NewMasterKeyService(saltStorage)
	masterKey, err := deriveMasterKey(saltStorage,"Please enter your master key")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	err = masterKeyService.Verify(masterKey, JsonStorage)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// then check the value
	secretValue := getSecretValue(args)

	// then encrypt the value an add it to the storage
	encrypter := crypt.NewAESEncrypter(masterKey)
	secret := domain.NewAddSecret(JsonStorage, encrypter)

	if err := secret.Execute(args[0], secretValue); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Secret added successfully 🔐")
}
