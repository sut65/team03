package entity

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name    string
	Price   int
	Item    int
	Service []Service `gorm:"foreignKey:FoodID"`
}

type Drink struct {
	gorm.Model
	Name    string
	Price   int
	Item    int
	Service []Service `gorm:"foreignKey:DrinkID"`
}

type Accessories struct {
	gorm.Model
	Name    string
	Price   int
	Item    int
	Service []Service `gorm:"foreignKey:AccessoriesID"`
}

// หลัก 1
type Service struct {
	gorm.Model
	CustomerID    *uint
	Customer      Customer `gorm:"references:id"`
	Time          time.Time
	FoodID        *uint
	Food          Food `gorm:"references:id"`
	DrinkID       *uint
	Drink         Drink `gorm:"references:id"`
	AccessoriesID *uint
	Accessories   Accessories `gorm:"references:id"`
}
