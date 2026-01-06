package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/Efesngl/learngo/todocli/storage"
	"github.com/Efesngl/learngo/todocli/types"
)

func Create(title string) error {
	list, err := Load()
	if err != nil {
		return err
	}
	todo := types.TodoItem{
		Id:     len(list) + 1,
		Title:  title,
		IsDone: false,
	}
	list = append(list, todo)
	err = storage.Save("todos.json", list)
	if err != nil {
		return err
	}
	return nil
}
func Mark(id int, state bool) error {
	list, err := Load()
	if err != nil {
		return err
	}
	for i := range list {
		if list[i].Id == id {
			list[i].IsDone = state
			break
		}
	}
	storage.Save("todos.json", list)
	return nil
}
func Delete(id int) error {
	list, err := Load()
	if err != nil {
		return err
	}

	for index, todo := range list {
		if todo.Id == id {
			newList := append(list[:index], list[index+1:]...)
			fmt.Println(newList)
			err := storage.Save("todos.json", newList)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("todo not found")
}
func Load() ([]types.TodoItem, error) {
	file, err := storage.Get("todos.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var TodoList = []types.TodoItem{}
	err = json.NewDecoder(file).Decode(&TodoList)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return TodoList, nil
		}
		return nil, err
	}
	return TodoList, nil
}
