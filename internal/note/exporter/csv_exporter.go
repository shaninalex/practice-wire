package exporter

import (
	"context"

	"github.com/shaninalex/practice-wire/internal/domain"
)

func NewCSVExporter(
	storage domain.IStorage,
) *CSVExporter {
	s := &CSVExporter{storage: storage}
	return s
}

type CSVExporter struct {
	storage domain.IStorage
}

func (s *CSVExporter) Export(ctx context.Context) (string, error) {
	return "", nil
}
