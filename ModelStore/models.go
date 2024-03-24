package modelstore

import (
	"gorm.io/gorm"
)

type CurrencyStore struct {
}

type UserStore struct {
	gorm.Model
	Fname string
	Lname string
	Phone string
	UpiID string
}

//Spends

type SpendStore struct {
	gorm.Model
	id      string
	GroupID string
}

// Groups

type GroupStore struct {
	gorm.Model
	id             string
	users          []UserStore
	homeCurrency   CurrencyStore
	foreignCurency CurrencyStore
}
