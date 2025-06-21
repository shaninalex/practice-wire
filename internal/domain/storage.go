package domain

import (
	"context"

	"github.com/google/uuid"
)

type IStorage interface {
	Save(ctx context.Context, note *Note) (*Note, error)
	Get(ctx context.Context, id uuid.UUID) (*Note, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, query string) ([]*Note, error)
}
