package exporter

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shaninalex/practice-wire/internal/domain"
)

func NewJSONExporter(
	storage domain.IStorage,
) *JSONExporter {
	s := &JSONExporter{
		storage: storage,
	}
	return s
}

type JSONExporter struct {
	storage domain.IStorage
}

func (s *JSONExporter) Export(ctx context.Context, destination string) (string, error) {
	notes, err := s.storage.List(ctx, "")
	if err != nil {
		return "", nil
	}
	b, err := json.Marshal(notes)
	if err != nil {
		return "", nil
	}
	path := filepath.Join(destination, "backup.json")
	if err := os.WriteFile(path, b, 0664); err != nil {
		return "", err
	}
	return path, nil
}
