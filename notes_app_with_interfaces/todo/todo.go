package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(todo)

	if err!= nil {
        return err
    }
	return os.WriteFile(filename, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
       return Todo{}, errors.New("Input cannot be empty")
    }  
	return Todo{
		Text: content,
	}, nil
}