package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Struct
type Todo struct {
	Text string `json:"text"`
}

//Struct tags are metadata that can be added to the struct fields(we use it here as the json package(encoding/json) can then come up with custom key in the JSON data corresponding to each struct field based on the tags that we provide in the aboce struct)

// Constructor
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	newNote := Todo{
		Text: content,
	}

	return newNote, nil
}

// Methods
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo) //for the note's fields to be accessible, the fields need to start with capital letters

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644) //it returns an error, so we're returning it. If it returns the error returned will be nil
}
