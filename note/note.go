package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdat time.Time
}

func (note Note) Display() {
	// Display the note
	fmt.Printf("Your note titled: %v has the following content:\n\n%v", note.title, note.content)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("input cannot be empty")
	}
	return Note{
		title:     title,
		content:   content,
		createdat: time.Now(),
	}, nil
}
