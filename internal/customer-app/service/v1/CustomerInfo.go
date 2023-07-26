package v1

import (
	"context"
	"github.com/PYxy/go-web/internal/customer-app/store"
)

type CustomerInfoSrv interface {
	Create(ctx context.Context, customer *store.Customer) error
	Update(ctx context.Context, customer *store.Customer) error
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) ([]store.Customer, error)
}

var _ CustomerInfoSrv = (*customerInfoService)(nil)

type customerInfoService struct {
	store store.Factory
}

func (c *customerInfoService) Create(ctx context.Context, customer *store.Customer) error {
	//TODO implement me
	return c.store.CustomerInfoOption().Create(ctx, customer)

}

func (c *customerInfoService) Update(ctx context.Context, customer *store.Customer) error {
	//TODO implement me
	return c.store.CustomerInfoOption().Update(ctx, customer)
}

func (c *customerInfoService) Delete(ctx context.Context, username string) error {
	//TODO implement me
	return c.store.CustomerInfoOption().Delete(ctx, username)
}

func (c *customerInfoService) Get(ctx context.Context, username string) ([]store.Customer, error) {
	//TODO implement me
	return c.store.CustomerInfoOption().Get(ctx, username)
}

func newcustomerINfoService(store store.Factory) *customerInfoService {
	return &customerInfoService{store: store}
}
