package models

import "github.com/jinzhu/gorm"

type Inventory struct {
	gorm.Model
	Items Items
}
