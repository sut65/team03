package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	G_Name   string
	Customer []Customer `gorm:"foreignKey:Gender_ID"`
}

type Province struct {
	gorm.Model
	P_Name string
	Customer []Customer `gorm:"foreignKey:Province_ID"`
}

type Nametitle struct {
	gorm.Model
	NT_Name string
	Customer    []Customer `gorm:"foreignKey:Nametitle_ID"`
}

type Customer struct {
	gorm.Model

	Nametitle_ID *uint
	Nametitle    Nametitle

	FirstName string
	LastName  string
	Password  string
	Age       uint
	Phone     string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`

	SigninID *uint
	Signin   Signin `gorm:"references:ID"`

	Gender_ID *uint
	Gender    Gender

	Province_ID *uint
	Province    Province



	Bookings  []Booking   `gorm:"foreignKey:CustomerID"`
	RepairReq []RepairReq `gorm:"foreignKey:CustomerID"`
	Review    []Review    `gorm:"foreignKey:CustomerID"`
	Service   []Service   `gorm:"foreignKey:CustomerID"`
	Payment   []Payment   `gorm:"foreignKey:CustomerID"`
}
