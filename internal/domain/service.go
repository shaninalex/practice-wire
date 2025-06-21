package domain

import (
	"context"
)

type INoteService interface {
	Save(ctx context.Context, title, content string) (*Note, error)
	Get(ctx context.Context, id int64) (*Note, error)
	Delete(ctx context.Context, id int64) error
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
