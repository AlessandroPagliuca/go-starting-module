package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	// Define the Note struct here with the following fields
	// with the json tags (struct tags) as shown in the test cases.
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	// Display the note
	fmt.Printf("Your note titled: %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	// Save the note to a file
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("input cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
