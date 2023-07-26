package v1

import (
	"context"
	"github.com/PYxy/go-web/internal/customer-app/store"
)

type CustomerGoodSrv interface {
	Create(ctx context.Context, customer *store.Good) error
	Update(ctx context.Context, customer *store.Good) error
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) ([]store.Good, error)
}
