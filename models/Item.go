// models/Item.go
package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model //grom จะสร้าง ID, CreatedAt, Description....
	Name       string
	Price      float64
}
