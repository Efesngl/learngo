package commands

import (
	"flag"
	"fmt"

	"github.com/Efesngl/learngo/passkey/crypt"
	"github.com/Efesngl/learngo/passkey/domain"
	"github.com/Efesngl/learngo/passkey/storage"
)

func Delete(args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	all := fs.Bool("all", false, "Deletes all secrets")
	force := fs.Bool("force", false, "Skip confirmation")
	if err := fs.Parse(args); err != nil {
		fmt.Println("Error on flag parsing: ", err)
	}
	// get secret store
	if *all {
		HandleDeleteAll(*force)
		return
	}
	if fs.NArg() == 0 {
		fmt.Println("Error: secret name required")
		return
	}
	handleDeleteOne(fs.Arg(0), *force)

}
func HandleDeleteAll(force bool) {
	saltStorage := storage.NewSaltStorage("salt.bin")
	JsonStorage := storage.NewJSONStorage("secrets.json")
	var masterKey []byte
	fmt.Print("Please enter master key: ")
	fmt.Scan(&masterKey)
	masterKeyService := crypt.NewMasterKeyService(saltStorage)
	err := masterKeyService.Verify(masterKey, JsonStorage)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !force {
		if !confirm("Do you really want to delete all secrets") {
			fmt.Println("Aborted!")
			return
		}
	}
	deleteAll := domain.NewDeleteAllSecret(JsonStorage)
	err = deleteAll.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All secrets deleted")
	return
}
func handleDeleteOne(name string, force bool) {
	saltStorage := storage.NewSaltStorage("salt.bin")
	JsonStorage := storage.NewJSONStorage("secrets.json")
	var masterKey []byte
	fmt.Print("Please enter master key: ")
	fmt.Scan(&masterKey)
	masterKeyService := crypt.NewMasterKeyService(saltStorage)
	err := masterKeyService.Verify(masterKey, JsonStorage)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !force {
		if !confirm(fmt.Sprint("Do you really want to delete ", name)) {
			fmt.Println("Aborted!")
			return
		}
	}

	secret := domain.NewDeleteSecret(JsonStorage)
	err = secret.Execute(name)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Secret deleted: ", name)
}
