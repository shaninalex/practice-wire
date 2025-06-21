package storage

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shaninalex/practice-wire/internal/domain"
)

const (
	fileDateLayout = "02-01-2006 15:04:05"
	fileHeaderSep  = "---"
)

// File based notes has this template.
//
// title: "Note Title"
// created_at: 18-10-2025 11:45:20
// updated_at: 19-10-2025 12:35:20
// ---
// content goes here

func makeFileNoteContent(note *domain.Note) ([]byte, error) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("title: \"%s\"\n", note.Title))
	b.WriteString(fmt.Sprintf("created_at: %s\n", note.CreatedAt.Format(fileDateLayout)))
	b.WriteString(fmt.Sprintf("updated_at: %s\n", note.UpdatedAt.Format(fileDateLayout)))
	b.WriteString(fileHeaderSep + "\n")
	b.WriteString(note.Content)

	return b.Bytes(), nil
}

func parseFileNote(data []byte, id uuid.UUID) (*domain.Note, error) {
	parts := strings.SplitN(string(data), fileHeaderSep, 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid note format: missing separator")
	}

	headers := strings.Split(strings.TrimSpace(parts[0]), "\n")
	content := strings.TrimSpace(parts[1])

	note := &domain.Note{
		ID:      id,
		Content: content,
	}

	for _, line := range headers {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "title:") {
			note.Title = strings.Trim(strings.TrimPrefix(line, "title:"), " \"")
		} else if strings.HasPrefix(line, "created_at:") {
			raw := strings.TrimSpace(strings.TrimPrefix(line, "created_at:"))
			t, err := time.Parse(fileDateLayout, raw)
			if err != nil {
				return nil, fmt.Errorf("invalid created_at: %w", err)
			}
			note.CreatedAt = t
		} else if strings.HasPrefix(line, "updated_at:") {
			raw := strings.TrimSpace(strings.TrimPrefix(line, "updated_at:"))
			t, err := time.Parse(fileDateLayout, raw)
			if err != nil {
				return nil, fmt.Errorf("invalid updated_at: %w", err)
			}
			note.UpdatedAt = t
		}
	}

	return note, nil
}
