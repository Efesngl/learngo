package main

import (
	"fmt"
	"os"

	"github.com/Efesngl/learngo/passkey/commands"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]
	switch command {
	case "add":
		commands.Add(args)
		return
	case "list":
		commands.List()
		return
	case "init":
		commands.Init()
		return
	case "get":
		commands.Get(args)
		return
	case "delete":
		commands.Delete(args)
		return
	default:
		printHelp()
		return
	}
}

func printHelp() {
	fmt.Println("Help: ")
}
