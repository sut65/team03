package entity

import (
	"time"

	"gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Name    string
	Payment []Payment `gorm:"foreignKey:PlaceID"`
}

type Method struct {
	gorm.Model
	Name            string
	Destination     string
	Picture         string
	PaymentMethodID int
	PaymentMethod   PaymentMethod `gorm:"references:ID"`
	Payment         []Payment     `gorm:"foreignKey:MethodID"`
}

type PaymentMethod struct {
	gorm.Model
	Name    string
	Method  []Method  `gorm:"foreignKey:PaymentMethodID"`
	Payment []Payment `gorm:"foreignKey:PaymentMethodID"`
}

type Payment struct {
	gorm.Model
	CustomerID      int           `valid:"required~Please Login"`
	Customer        Customer      `valid:"-" gorm:"references:id"`
	PaymentMethodID int           `valid:"required~Choose payment method"`
	PaymentMethod   PaymentMethod `gorm:"references:ID"`
	MethodID        int           `valid:"required~Choose method"`
	Method          Method        `gorm:"references:ID"`
	PlaceID         int           `valid:"required~Choose place"`
	Place           Place         `gorm:"references:ID"`
	Time            time.Time
	Picture         string
}
