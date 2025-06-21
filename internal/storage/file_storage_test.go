package storage

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shaninalex/practice-wire/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_FileStorage_SaveAndGet(t *testing.T) {
	ctx := context.Background()
	tmpDir := t.TempDir()
	storage := NewFileStorage(tmpDir)

	note := &domain.Note{ID: uuid.New(), Title: "Hello", Content: "World"}
	_, err := storage.Save(ctx, note)
	assert.NoError(t, err)

	saved, err := storage.Get(ctx, note.ID)
	assert.NoError(t, err)
	assert.Equal(t, note.Content, saved.Content)
	assert.Equal(t, note.Title, saved.Title)
}

func Test_FileStorage_Remove(t *testing.T) {
	ctx := context.Background()
	tmpDir := t.TempDir()
	storage := NewFileStorage(tmpDir)

	note := &domain.Note{ID: uuid.New(), Title: "Hello", Content: "World"}
	_, _ = storage.Save(ctx, note)

	err := storage.Delete(ctx, note.ID)
	assert.NoError(t, err)

	_, err = storage.Get(ctx, note.ID)
	assert.Error(t, err)
}

func Test_FileStorage_List(t *testing.T) {
	ctx := context.Background()
	tmpDir := t.TempDir()
	storage := NewFileStorage(tmpDir)

	expected := map[uuid.UUID]*domain.Note{
		uuid.New(): {Title: "Hello 1", Content: "World 1"},
		uuid.New(): {Title: "Hello 2", Content: "World 2"},
	}

	// Assign IDs to notes and save them
	for id, note := range expected {
		note.ID = id
		_, err := storage.Save(ctx, note)
		assert.NoError(t, err)
	}

	savedNotes, err := storage.List(ctx, "")
	assert.NoError(t, err)
	assert.Len(t, savedNotes, len(expected))

	for _, sn := range savedNotes {
		expectedNote, ok := expected[sn.ID]
		assert.True(t, ok, "unexpected note with ID %s", sn.ID)
		assert.Equal(t, expectedNote.Title, sn.Title)
		assert.Equal(t, expectedNote.Content, sn.Content)
	}
}
