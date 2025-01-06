package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

// type Outputtable interface {
// 	Save() error
// 	Display()
// }

// type Outputtable interface {
// 	Saver
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {

	err := data.Save()

	if err != nil {
		return err
	}

	fmt.Println("saved successfully")

	return nil
}

func getNoteData() (string, string) {

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {

	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
