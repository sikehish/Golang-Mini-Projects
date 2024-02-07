package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// Constrcutor
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	newNote := Note{
		title: title, content: content, createdAt: time.Now(),
	}

	return newNote, nil
}

// Method
func (note Note) Display() {
	fmt.Printf("\nYour note titled %v has the following content:\n\n%v\nCreated At: %v\n\n", note.title, note.content, note.createdAt)
}
