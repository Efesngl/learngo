package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Efesngl/learngo/todocli/types"
)

func Get(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Save(fileName string, todoList []types.TodoItem) error {
	fmt.Printf("Storage list %v\n", todoList)
	file, err := Get(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(todoList)
	if err != nil {
		return err
	}
	fmt.Printf("Json data: %v\n", string(jsonData))
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.WriteAt(jsonData, 0)
	if err != nil {
		return err
	}
	return nil
}
