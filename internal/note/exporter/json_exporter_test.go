package exporter_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/shaninalex/practice-wire/internal/domain"
	"github.com/shaninalex/practice-wire/internal/note/exporter"
	"github.com/shaninalex/practice-wire/internal/storage"
	"github.com/stretchr/testify/assert"
)

func Test_JsonExporter(t *testing.T) {
	ctx := context.Background()
	tmpFile := filepath.Join(t.TempDir(), "test.db")
	storage, err := storage.NewSqliteStorage(tmpFile)
	assert.NoError(t, err)

	noteA := &domain.Note{ID: uuid.New(), Title: "Hello", Content: "World"}
	_, _ = storage.Save(ctx, noteA)

	noteB := &domain.Note{ID: uuid.New(), Title: "Hello", Content: "World"}
	_, _ = storage.Save(ctx, noteB)

	jsonExporter, err := exporter.ExporterFactory(domain.ExportFormatJSON, storage)
	assert.NoError(t, err)

	destination := t.TempDir()
	backupPath, err := jsonExporter.Export(ctx, destination)
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(destination, "backup.json"), backupPath)

	b, err := os.ReadFile(backupPath)
	assert.NoError(t, err)

	backupNotes := []*domain.Note{}
	err = json.Unmarshal(b, &backupNotes)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(backupNotes))

	assert.Equal(t, backupNotes[0].Content, noteA.Content)
	assert.Equal(t, backupNotes[0].Title, noteA.Title)
	assert.Equal(t, backupNotes[0].ID, noteA.ID)

	assert.Equal(t, backupNotes[1].Content, noteB.Content)
	assert.Equal(t, backupNotes[1].Title, noteB.Title)
	assert.Equal(t, backupNotes[1].ID, noteB.ID)
}
