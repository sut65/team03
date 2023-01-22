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

type Crypto struct {
	gorm.Model
	Name            string
	PublicKey       string
	Picture         string
	PaymentMethodID int
	PaymentMethod   PaymentMethod `gorm:"references:ID"`
}

type Bank struct {
	gorm.Model
	Name            string
	Number          string
	Picture         string
	PaymentMethodID int
	PaymentMethod   PaymentMethod `gorm:"references:ID"`
}

type PaymentMethod struct {
	gorm.Model
	Name    string
	Bank    []Bank    `gorm:"foreignKey:PaymentMethodID"`
	Crypto  []Crypto  `gorm:"foreignKey:PaymentMethodID"`
	Payment []Payment `gorm:"foreignKey:PaymentMethodID"`
}

type Payment struct {
	gorm.Model
	CustomerID      *uint
	Customer        Customer `gorm:"references:ID"`
	PaymentMethodID *uint
	PaymentMethod   PaymentMethod `gorm:"references:ID"`
	PlaceID         *uint
	Place           Place `gorm:"references:ID"`
	Time            time.Time
	Picture         string
}
