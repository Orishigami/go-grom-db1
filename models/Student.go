// models/Student.go
package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model //grom จะสร้าง ID, CreatedAt, Description....
	Name       string
	Age        int
}
