package domain

import "context"

type IStorage interface {
	Save(ctx context.Context, note *Note) (*Note, error)
	Get(ctx context.Context, id int64) (*Note, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, query string) ([]*Note, error)
}
