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
	CustomerID      int      `valid:"required~Please Login"`
	Customer        Customer `gorm:"references:id"`
	Time            time.Time
	FoodID          int         `valid:"required~Choose Food"`
	Food            Food        `gorm:"references:id"`
	FoodItem        int         `valid:"range(0|50)"`
	DrinkID         int         `valid:"required~Choose Drink"`
	Drink           Drink       `gorm:"references:id"`
	DrinkItem       int         `valid:"range(0|50)"`
	AccessoriesID   int         `valid:"required~Choose Accessories"`
	Accessories     Accessories `gorm:"references:id"`
	AccessoriesItem int         `valid:"range(0|50)"`
}
