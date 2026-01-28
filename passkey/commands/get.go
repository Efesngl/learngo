package commands

import (
	"fmt"
	"github.com/Efesngl/learngo/passkey/crypt"
	"github.com/Efesngl/learngo/passkey/domain"
	"github.com/Efesngl/learngo/passkey/storage"
)

func Get(args []string) {
	store := storage.NewJSONStorage("secrets.json")

	saltStorage := storage.NewSaltStorage("salt.bin")
	MasterKeyService := crypt.NewMasterKeyService(saltStorage)

	var password []byte
	fmt.Print("Please enter master password: ")
	fmt.Scan(&password)
	
	masterKey, err := MasterKeyService.Derive(password)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	encrypter := crypt.NewAESEncrypter(masterKey)
	secret := domain.NewGetSecret(store, encrypter)

	value, err := secret.Execute(args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Password: ", value)
}
