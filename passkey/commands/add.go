package commands

import (
	"fmt"

	"github.com/Efesngl/learngo/passkey/crypt"
	"github.com/Efesngl/learngo/passkey/domain"
	"github.com/Efesngl/learngo/passkey/storage"
)

func Add(args []string) {
	// get secret store
	store := storage.NewJSONStorage("secrets.json")

	// get salt
	saltStorage := storage.NewSaltStorage("salt.bin")

	// get masterkey
	var password []byte
	fmt.Print("Please enter master password: ")
	fmt.Scan(&password)
	MasterKeyService := crypt.NewMasterKeyService(saltStorage)
	masterKey, err := MasterKeyService.Derive(password)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	encrypter := crypt.NewAESEncrypter(masterKey)
	secret := domain.NewAddSecret(store, encrypter)

	if err := secret.Execute(args[0], args[1]); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Secret added successfully 🔐")
}
