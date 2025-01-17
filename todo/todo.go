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
	// Display the todo
	fmt.Printf("Your todo text: \n\n%v\n\n", todo.Text)
}

func (todo Todo) Save() error {
	// Save the todo to a file
	fileName := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("input cannot be empty")
	}
	return Todo{
		Text: content,
	}, nil
}
