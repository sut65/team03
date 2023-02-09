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

// หลัก 1.
type Service struct {
	gorm.Model
	CustomerID    int      `valid:"-"`
	Customer      Customer `gorm:"references:id"`
	Time          time.Time
	FoodID        int         `valid:"required~Please Choose some order of Food"`
	Food          Food        `gorm:"references:id"`
	DrinkID       int         `valid:"required~Please Choose some order of Drink"`
	Drink         Drink       `gorm:"references:id"`
	AccessoriesID int         `valid:"required~Please Choose some order of Accessories"`
	Accessories   Accessories `gorm:"references:id"`
}
