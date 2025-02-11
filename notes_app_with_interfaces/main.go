package main

import (
	"bufio"
	"fmt"
	"notes_app/notes"
	"strings"
	"os"
	"notes_app/todo"
)

type saver interface {
	Save() error
}

type Displayer interface {
	saver
	Display()
}

func outputData(data Displayer) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err!= nil {
		fmt.Println("Error saving data")
        return err
	}
	fmt.Println("Data saved successfully")
	return nil
}

func getTodoData() string {
	text := getUserInput("Todo Text: ")
	return text
}

func getNoteData() (string, string){
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err!= nil {
        fmt.Println("Error reading input", err)
        return ""
    }

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func main() {
	title, content := getNoteData()

	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
        return
    }

	UserNote, err := note.New(title, content)

	if err!= nil {
        fmt.Println(err)
        return
    }
    
	err = outputData(UserNote)

	if err != nil {
        return
	}
	fmt.Println("Note saved successfully")

	err = outputData(todo)

	if err!= nil {

        return
    }

	fmt.Println("Todo saved successfully")
}
