package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Gender   string
	Customer []Customer `gorm:"foreignKey:Gender_ID"`
}

type Province struct {
	gorm.Model
	Province string
	Customer []Customer `gorm:"foreignKey:Province_ID"`
}

type Memberlevel struct {
	gorm.Model
	Memberlevel string
	Customer    []Customer `gorm:"foreignKey:Memberlevel_ID"`
}

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Password  string
	Age       uint
	Phone     string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`

	Gender_ID *uint
	Gender    Gender

	Province_ID *uint
	Province    Province

	Memberlevel_ID *uint
	Memberlevel    Memberlevel

	Bookings []Booking `gorm:"foreignKey:CustomerID"`

	Review []Review `gorm:"foreignKey:DepartmentID"`
}
