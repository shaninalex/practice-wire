package exporter

import (
	"context"

	"github.com/shaninalex/practice-wire/internal/domain"
)

func NewXMLExporter(
	storage domain.IStorage,
) *XMLExporter {
	s := &XMLExporter{
		storage: storage,
	}
	return s
}

type XMLExporter struct {
	storage domain.IStorage
}

func (s *XMLExporter) Export(ctx context.Context) (string, error) {
	return "", nil
}
