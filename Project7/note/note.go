package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct
type Note struct {
	Title     string    `json:"title"` //-> struct tag
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

//Struct tags are metadata that can be added to the struct fields(we use it here as the json package(encoding/json) can then come up with custom key in the JSON data corresponding to each struct field based on the tags that we provide in the aboce struct)

// Constructor
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	newNote := Note{
		Title: title, Content: content, CreatedAt: time.Now(),
	}

	return newNote, nil
}

// Methods
func (note Note) Display() {
	fmt.Printf("\nYour note titled %v has the following content:\n\n%v\nCreated At: %v\n\n", note.Title, note.Content, note.CreatedAt)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(note) //for the note's fields to be accessible, the fields need to start with capital letters

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644) //it returns an error, so we're returning it. If it returns the error returned will be nil
}
