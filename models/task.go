package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:name`
	Description string `json:description`
}
