package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/shaninalex/practice-wire/internal/domain"
)

type FileStorage struct {
	root string
}

func NewFileStorage(path string) *FileStorage {
	s := &FileStorage{
		root: path,
	}
	s.init()
	return s
}

func (s *FileStorage) init() {
}

func (s *FileStorage) Save(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	if note.ID == uuid.Nil {
		note.ID = uuid.New()
	}

	if err := os.MkdirAll(s.root, 0755); err != nil {
		return nil, fmt.Errorf("failed to ensure root dir: %w", err)
	}

	filename := fmt.Sprintf("%s.md", note.ID.String())
	path := filepath.Join(s.root, filename)

	content, err := makeFileNoteContent(note)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, content, 0664); err != nil {
		return nil, err
	}
	return note, nil
}

func (s *FileStorage) Get(ctx context.Context, id uuid.UUID) (*domain.Note, error) {
	entries, err := os.ReadDir(s.root)
	if err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("%s.md", id.String())

	for _, file := range entries {
		if file.Name() == filename {
			path := filepath.Join(s.root, filename)
			b, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("unable to open note %s", id.String())
			}

			note, err := parseFileNote(b, id)
			if err != nil {
				return nil, err
			}
			return note, nil
		}
	}
	return nil, fmt.Errorf("note with id %s was not found", id.String())
}

func (s *FileStorage) Delete(ctx context.Context, id uuid.UUID) error {
	entries, err := os.ReadDir(s.root)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.md", id.String())
	for _, file := range entries {
		if file.Name() == filename {
			if err := os.Remove(filepath.Join(s.root, filename)); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("note with id %s was not found", id.String())
}

func (s *FileStorage) List(ctx context.Context, query string) ([]*domain.Note, error) {
	entries, err := os.ReadDir(s.root)
	if err != nil {
		return nil, err
	}

	notes := make([]*domain.Note, 0)

	for _, file := range entries {
		path := filepath.Join(s.root, file.Name())
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("unable to open note %s", file.Name())
		}

		id := uuid.MustParse(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
		note, err := parseFileNote(b, id)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
