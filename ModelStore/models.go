package modelstore

import (
	"gorm.io/gorm"
)

type UserStore struct {
	gorm.Model
	Fname string
	Lname string
	Phone string
	UpiID string
}
