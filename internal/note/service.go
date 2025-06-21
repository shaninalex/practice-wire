package note

import (
	"context"
	"github.com/shaninalex/practice-wire/internal/domain"
	"time"
)

type NoteService struct {
	storage domain.IStorage
}

func (s *NoteService) Save(ctx context.Context, title, content string) (*domain.Note, error) {
	note := domain.Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return s.storage.Save(ctx, &note)
}

func (s *NoteService) Get(ctx context.Context, id int64) (*domain.Note, error) {
	return s.storage.Get(ctx, id)
}

func (s *NoteService) Delete(ctx context.Context, id int64) error {
	return s.storage.Delete(ctx, id)
}

func (s *NoteService) List(ctx context.Context, query string) ([]*domain.Note, error) {
	return s.storage.List(ctx, query)
}

func (s *NoteService) Export(ctx context.Context, format domain.ExportFormat, destination string) error {
	panic("implement me")
}
