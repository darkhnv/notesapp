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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Title:\n\n%v\n\nContent:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	data, err := json.MarshalIndent(note, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("both title and content are required")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
