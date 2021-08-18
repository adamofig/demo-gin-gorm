package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Company  string `json:"company"`
	UserType string `json:"userType"`
}
