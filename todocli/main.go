package main

import (
	"fmt"
	"github.com/Efesngl/learngo/todocli/commands"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a command")
		return
	}
	fmt.Println(args[2:])

	command, ok := commands.CommandsMap[args[1]]
	if !ok {
		log.Fatal("Command not found")
	}
	// fmt.Println(args[2:])
	command.Run(args[2:])
}
