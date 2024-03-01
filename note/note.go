package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note represents a note
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Display displays the note
func (note Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\n", note.Title, note.Content)
}

// Save saves the note to a JSON file
func (note Note) Save() error {
	// Generate file name from note title
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"

	// Convert note to JSON
	jsonData, err := json.Marshal(note)
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

// New creates a new note with the given title, content, and current timestamp
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, fmt.Errorf("both title and content are required")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
