package store

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

//定义数据实例对象

// Customer  用户信息
type Customer struct {
	Status int `json:"status" gorm:"column:status" validate:"omitempty"`

	Name string `json:"name" gorm:"column:name" validate:"required,min=1,max=30"`

	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`

	// Required: true
	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	TotalPolicy int64 `json:"totalPolicy" gorm:"-" validate:"omitempty"`

	HobbySlice []string `json:"hobbySlice" gorm:"-" validate:"omitempty"`

	Hobby string `json:"hobby" gorm:"hobby" validate:"omitempty"`
}

func (u *Customer) TableName() string {
	return "customer"
}

// BeforeSave  插入记录之后更新某个字段 Update
func (u *Customer) BeforeSave(tx *gorm.DB) error {
	fmt.Println("AfterCreate")
	if u.HobbySlice == nil || len(u.HobbySlice) == 0 {
		u.Hobby = ""
	} else {
		if bytes, err := json.Marshal(u.HobbySlice); err != nil {
			return err
		} else {
			u.Hobby = string(bytes)
		}
	}
	return nil
}

func (u *Customer) BeforeUpdate(tx *gorm.DB) error {
	if u.HobbySlice == nil || len(u.HobbySlice) == 0 {
		u.Hobby = ""
	} else {
		if bytes, err := json.Marshal(u.HobbySlice); err != nil {
			return err
		} else {
			u.Hobby = string(bytes)
		}
	}
	return nil
}

///根据需求可以添加更多的方法

// Good 商品信息
type Good struct {
}
