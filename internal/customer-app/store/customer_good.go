package store

import "context"

// CustomerGoodStore defines the Customer storage interface.
type CustomerGoodStore interface {
	Set(ctx context.Context, username string, good *Good) error
	Update(ctx context.Context, username string, good *Good) error
	Delete(ctx context.Context, username string, goodName *Good) error
	Get(ctx context.Context, username string) (*Good, error)
}
