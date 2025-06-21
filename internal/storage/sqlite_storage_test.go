package storage

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/shaninalex/practice-wire/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_SqliteStorage_CRUD(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test.db")
	storage, err := NewSqliteStorage(tmpFile)
	assert.NoError(t, err)

	note := &domain.Note{Title: "Hello", Content: "World"}
	saved, err := storage.Save(context.Background(), note)
	assert.NoError(t, err)
	assert.NotZero(t, saved.ID)

	got, err := storage.Get(context.Background(), saved.ID)
	assert.NoError(t, err)
	assert.Equal(t, saved.Title, got.Title)
}

func Test_SqliteStorage_Remove(t *testing.T) {
	ctx := context.Background()
	tmpFile := filepath.Join(t.TempDir(), "test.db")
	storage, err := NewSqliteStorage(tmpFile)
	assert.NoError(t, err)

	note := &domain.Note{ID: uuid.New(), Title: "Hello", Content: "World"}
	_, _ = storage.Save(ctx, note)

	err = storage.Delete(ctx, note.ID)
	assert.NoError(t, err)

	_, err = storage.Get(ctx, note.ID)
	assert.Error(t, err)
}

func Test_SqliteStorage_List(t *testing.T) {
	ctx := context.Background()
	tmpFile := filepath.Join(t.TempDir(), "test.db")
	storage, err := NewSqliteStorage(tmpFile)
	assert.NoError(t, err)

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
		// assert.Equal(t, expectedNote.Title, sn.Title)
		assert.Equal(t, expectedNote.Content, sn.Content)
	}
}
