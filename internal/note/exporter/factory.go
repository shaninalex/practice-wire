// internal/note/exporter/factory.go
package exporter

import (
	"fmt"

	"github.com/shaninalex/practice-wire/internal/domain"
)

func ExporterFactory(format domain.ExportFormat, storage domain.IStorage) (domain.IExporter, error) {
	switch format {
	case domain.ExportFormatJSON:
		return ProvideJSONExporter(storage), nil
	case domain.ExportFormatCSV:
		return ProvideCSVExporter(storage), nil
	case domain.ExportFormatXML:
		return ProvideXMLExporter(storage), nil
	case domain.ExportFormatMarkdown:
		return ProvideMarkdownExporter(storage), nil

	default:
		return nil, fmt.Errorf("unsupported export format: %s", format)
	}
}
