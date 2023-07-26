package v1

import "github.com/PYxy/go-web/internal/customer-app/store"

type Service interface {
	CustomerInfoSrv() CustomerInfoSrv
	CustomerGoodOption() CustomerGoodSrv
}

type service struct {
	store store.Factory
}

func (s *service) CustomerInfoSrv() CustomerInfoSrv {
	//TODO implement me
	return newcustomerINfoService(s.store)

}

func (s *service) CustomerGoodOption() CustomerGoodSrv {
	//TODO implement me
	panic("implement me")
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
