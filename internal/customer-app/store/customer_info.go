package store

import "context"

// CustomerInfoStore defines the Customer storage interface.
type CustomerInfoStore interface {
	Create(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) ([]Customer, error)
}
