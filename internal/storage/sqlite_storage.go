package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shaninalex/practice-wire/internal/domain"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(path string) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite DB: %w", err)
	}

	s := &SqliteStorage{db: db}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SqliteStorage) init() error {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *SqliteStorage) Save(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	now := time.Now()

	if note.ID == uuid.Nil {
		return s.insertNewNote(ctx, note, now)
	}

	note.UpdatedAt = now
	updated, err := s.updateExistingNote(ctx, note)
	if err != nil {
		return nil, err
	}

	if updated {
		return note, nil
	}

	return s.insertNoteWithID(ctx, note, now)
}

func (s *SqliteStorage) Get(ctx context.Context, id uuid.UUID) (*domain.Note, error) {
	row := s.db.QueryRowContext(ctx,
		`SELECT id, title, content, created_at, updated_at FROM notes WHERE id = ?`,
		id.String())

	var note domain.Note
	var idStr string
	if err := row.Scan(&idStr, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("note not found")
		}
		return nil, err
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID in DB: %w", err)
	}
	note.ID = parsedID

	return &note, nil
}

func (s *SqliteStorage) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM notes WHERE id = ?`, id.String())
	return err
}

func (s *SqliteStorage) List(ctx context.Context, query string) ([]*domain.Note, error) {
	var rows *sql.Rows
	var err error

	if query == "" {
		rows, err = s.db.QueryContext(ctx, `SELECT id, title, content, created_at, updated_at FROM notes`)
	} else {
		rows, err = s.db.QueryContext(ctx,
			`SELECT id, title, content, created_at, updated_at FROM notes WHERE title LIKE ?`,
			"%"+query+"%")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*domain.Note
	for rows.Next() {
		var note domain.Note
		var idStr string
		if err := rows.Scan(&idStr, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		note.ID, err = uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("invalid UUID in DB: %w", err)
		}
		notes = append(notes, &note)
	}

	return notes, nil
}

func (s *SqliteStorage) insertNewNote(ctx context.Context, note *domain.Note, now time.Time) (*domain.Note, error) {
	note.ID = uuid.New()
	note.CreatedAt = now
	note.UpdatedAt = now

	_, err := s.db.ExecContext(ctx, `
		INSERT INTO notes (id, title, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`,
		note.ID.String(), note.Title, note.Content, note.CreatedAt, note.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *SqliteStorage) updateExistingNote(ctx context.Context, note *domain.Note) (bool, error) {
	res, err := s.db.ExecContext(ctx, `
		UPDATE notes SET title = ?, content = ?, updated_at = ? WHERE id = ?`,
		note.Title, note.Content, note.UpdatedAt, note.ID.String())
	if err != nil {
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}

func (s *SqliteStorage) insertNoteWithID(ctx context.Context, note *domain.Note, now time.Time) (*domain.Note, error) {
	note.CreatedAt = now
	note.UpdatedAt = now

	_, err := s.db.ExecContext(ctx, `
		INSERT INTO notes (id, title, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`,
		note.ID.String(), note.Title, note.Content, note.CreatedAt, note.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return note, nil
}
