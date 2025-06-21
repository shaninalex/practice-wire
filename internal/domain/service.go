package domain

import (
	"context"

	"github.com/google/uuid"
)

type INoteService interface {
	Save(ctx context.Context, title, content string) (*Note, error)
	Get(ctx context.Context, id uuid.UUID) (*Note, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, query string) ([]*Note, error)
	Export(cxt context.Context, format ExportFormat, destination string) error
}

type ExportFormat string

const (
	ExportFormatMarkdown ExportFormat = "markdown"
	ExportFormatCSV      ExportFormat = "csv"
	ExportFormatJSON     ExportFormat = "json"
	ExportFormatXML      ExportFormat = "xml"
)
