package exporter

import (
	"context"

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

func (s *JSONExporter) Export(ctx context.Context) (string, error) {
	return "", nil
}
