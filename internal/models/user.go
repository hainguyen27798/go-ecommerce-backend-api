package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:user_name; unique;"`
	IsActive bool   `gorm:"column:is_active; type:boolean;"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

func (user *User) TableName() string {
	return "users"
}
