package domain

import "context"

// IExporter describe exporter interface
type IExporter interface {
	// Export create a single file with containing all notes
	// return created file path or an error
	Export(ctx context.Context) (string, error)
}
