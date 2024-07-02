package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"column:name; unique;"`
	Note  bool   `gorm:"column:note; type:text;"`
	Users []User `gorm:"many2many:user_roles;"`
}

func (role *Role) TableName() string {
	return "roles"
}
