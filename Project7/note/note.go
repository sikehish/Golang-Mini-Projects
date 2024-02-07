package note

import (
	"errors"
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
