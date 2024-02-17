package main

import (
	"bufio"
	"fmt"
	"notesapp/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	if err := userNote.Save(); err != nil {
		fmt.Println("Failed to save the note")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	// removing the new line
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
