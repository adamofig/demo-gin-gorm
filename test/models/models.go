package models

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	Name string
}
