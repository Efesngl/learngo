package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Efesngl/learngo/todocli/todos"
	"github.com/Efesngl/learngo/todocli/types"
)

var CommandsMap map[string]types.Command

func init() {
	CommandsMap = map[string]types.Command{
		"help": {Name: "help", Run: Help, Description: "Show help", Hint: "help"},
		"add":  {Name: "add", Run: Add, Description: "Add a todo", Hint: "add <todo>"},
		"list": {Name: "list", Run: List, Description: "List todos", Hint: "list"},
		"mark": {Name: "mark", Run: Mark, Description: "Mark a todo", Hint: "mark <todo> <state 1|0>"},
		"del":  {Name: "del", Run: Delete, Description: "Delete a todo", Hint: "del <todo>"},
	}

}

func Help(args []string) {
	for _, command := range CommandsMap {
		fmt.Printf("%v: %v, %v\n", command.Name, command.Hint, command.Description)
	}
}
func Add(args []string) {
	str := strings.Join(args, " ")
	todos.Create(str)
}
func List(args []string) {
	list, err := todos.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID\tTitle\tIs Done")
	for _, todo := range list {
		fmt.Printf("%v\t%v\t%v\n", todo.Id, todo.Title, todo.IsDone)
	}
}
func Mark(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	state := true
	if args[1] == "0" {
		state = false
	}
	todos.Mark(id, state)
	fmt.Printf("marked %v as %v\n", id, state)
}
func Delete(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	err = todos.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("todo deleted")
}
