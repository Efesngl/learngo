package commands

import (
	"fmt"

	"github.com/Efesngl/learngo/passkey/crypt"
	"github.com/Efesngl/learngo/passkey/domain"
)

func ChangeMasterKey(args []string) {
	saltStore, JsonStore := initStorages("salt.bin", "secrets.json")
	masterKeyService := crypt.NewMasterKeyService(saltStore)
	oldMasterKey, err := deriveMasterKey(saltStore, "Please enter your old master key: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	err = masterKeyService.Verify(oldMasterKey, JsonStore)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	newMasterKey, err := deriveMasterKey(saltStore, "Please enter your new master key: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	oldCrypt := crypt.NewAESEncrypter(oldMasterKey)
	newCrypt := crypt.NewAESEncrypter(newMasterKey)
	changeMasterKey := domain.NewChangeMasterKey(JsonStore, oldCrypt, newCrypt)
	err = changeMasterKey.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Master key changed successfully")
}
