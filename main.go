package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"notesapp/note"
	"notesapp/todo"
)

// Define interfaces for types that can save data and display data
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	// Get input for note title and content
	title, content := getNoteData()

	// Get input for todo item
	todoText := getUserInput("Enter todo text: ")

	// Create todo item and handle error
	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create note and handle error
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display and save todo item
	err = outputData(todoItem)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display and save note
	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// outputData displays the data and saves it
func outputData(data outputtable) error {
	data.Display() // Display the data
	return saveData(data)
}

// saveData saves the data
func saveData(data saver) error {
	// Attempt to save the data and handle error
	if err := data.Save(); err != nil {
		return fmt.Errorf("failed to save data: %v", err)
	}

	fmt.Println("Saving data succeeded!")
	return nil
}

// getNoteData gets input for note title and content
func getNoteData() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")
	return title, content
}

// getUserInput prompts the user for input and returns it
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Trim whitespace and newline characters
	text = strings.TrimSpace(text)

	return text
}
