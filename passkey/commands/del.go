package commands

import (
	"fmt"
	"github.com/Efesngl/learngo/passkey/domain"
	"github.com/Efesngl/learngo/passkey/storage"
)

func Delete(name string) {
	// get secret store
	store := storage.NewJSONStorage("secrets.json")
	secret := domain.NewDeleteSecret(store)
	err := secret.Execute(name)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Secret deleted: ", name)
}
