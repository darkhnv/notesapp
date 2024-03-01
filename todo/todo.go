package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo represents a todo item
type Todo struct {
	Text string `json:"text"`
}

// Display displays the todo item
func (todo Todo) Display() {
	fmt.Println("Todo:", todo.Text)
}

// Save saves the todo item to a JSON file
func (todo Todo) Save() error {
	fileName := "todo.json"

	// Convert todo item to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Write JSON data to file
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

// New creates a new todo item with the given text
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, fmt.Errorf("todo text cannot be empty")
	}
	return Todo{Text: text}, nil
}
