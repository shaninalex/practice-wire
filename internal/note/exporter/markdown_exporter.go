package exporter

import (
	"context"

	"github.com/shaninalex/practice-wire/internal/domain"
)

func NewMarkdownExporter(
	storage domain.IStorage,
) *MarkdownExporter {
	s := &MarkdownExporter{
		storage: storage,
	}
	return s
}

type MarkdownExporter struct {
	storage domain.IStorage
}

func (s *MarkdownExporter) Export(ctx context.Context) (string, error) {
	return "", nil
}
