package commands

import (
	"fmt"
	"github.com/Efesngl/learngo/passkey/domain"
	"github.com/Efesngl/learngo/passkey/storage"
)

func List() {
	store := storage.NewJSONStorage("secrets.json")
	listSecrets := domain.NewListSecrets(store)
	secrets, err := listSecrets.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, s := range secrets {
		fmt.Printf("%s\n", s.Name)
	}

}
