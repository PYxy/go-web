package mysql

import (
	"context"
	"fmt"
	"github.com/PYxy/go-web/internal/customer-app/store"
	"gorm.io/gorm"
)

type customerInfo struct {
	db *gorm.DB
}

var _ store.CustomerInfoStore = (*customerInfo)(nil)

// Create 创建单个用户
// https://gorm.io/zh_CN/docs/create.html
func (c *customerInfo) Create(ctx context.Context, customer *store.Customer) error {
	//TODO implement me
	fmt.Println("插入数据..")

	return c.db.WithContext(ctx).Create(customer).Error
}

// Update 更新单个用户信息
// https://gorm.io/zh_CN/docs/update.html
func (c *customerInfo) Update(ctx context.Context, customer *store.Customer) error {
	//TODO implement me
	return c.db.WithContext(ctx).Save(customer).Error
}

// Delete 根据用户名删除用户
// https://gorm.io/zh_CN/docs/delete.html
func (c *customerInfo) Delete(ctx context.Context, username string) error {
	//TODO implement me
	return c.db.WithContext(ctx).Where("name = ?", username).Delete(&store.Customer{}).Error

}

// Get 根据用户名获取用户信息
func (c *customerInfo) Get(ctx context.Context, username string) ([]store.Customer, error) {
	//TODO implement me
	var resultCustomer []store.Customer
	return resultCustomer, c.db.WithContext(ctx).Where("name = ?", username).Find(&resultCustomer).Error
}

func newCustomerInfo(ds *mysqlStore) *customerInfo {
	return &customerInfo{db: ds.db}
}
