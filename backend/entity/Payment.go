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
	CustomerID      *uint
	Customer        Customer `gorm:"references:ID"`
	PaymentMethodID *uint
	PaymentMethod   PaymentMethod `gorm:"references:ID"`
	MethodID        *uint
	Method          Method `gorm:"references:ID"`
	PlaceID         *uint
	Place           Place `gorm:"references:ID"`
	Time            time.Time
	Picture         string
}
