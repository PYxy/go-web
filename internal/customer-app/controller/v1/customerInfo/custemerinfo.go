package customerInfo

import (
	v1 "github.com/PYxy/go-web/internal/customer-app/service/v1"
	"github.com/PYxy/go-web/internal/customer-app/store"
)

type InfoController struct {
	srv v1.Service
}

func NewCustomerInfoController(store store.Factory) *InfoController {
	return &InfoController{
		srv: v1.NewService(store),
	}
}
