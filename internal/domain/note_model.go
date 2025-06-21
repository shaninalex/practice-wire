package domain

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (note *Note) WithID(iD uuid.UUID) *Note {
	note.ID = iD
	return note
}

func (note *Note) WithTitle(title string) *Note {
	note.Title = title
	return note
}

func (note *Note) WithContent(content string) *Note {
	note.Content = content
	return note
}

func (note *Note) WithCreatedAt(createdAt time.Time) *Note {
	note.CreatedAt = createdAt
	return note
}

func (note *Note) WithUpdatedAt(updatedAt time.Time) *Note {
	note.UpdatedAt = updatedAt
	return note
}

// Parse - content string in format:
//
// title: "Note Title"
// created_at: 18-10-2025 11:45:20
// updated_at: 19-10-2025 12:35:20
// ---
// content goes here
func (note *Note) Parse(fileContent string) *Note {
	panic("implement me")
}

func (note *Note) Build() Note {
	return *note
}

func DefaultNote() *Note {
	return &Note{}
}
