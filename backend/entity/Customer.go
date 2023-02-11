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
 
	Nametitle_ID *uint `valid:"required~Please select Nametitle"`
	Nametitle    Nametitle `valid:"-" gorm:"references:id"`

	FirstName string `valid:"required~FirstName not blank"`
	LastName  string `valid:"required~LastName not blank"`
	Password  string	
	Age      	int   `valid:"required~กรุณาระบุอายุ, range(0|150)~อายุไม่สามารถติดลบได้"` 
	Phone     string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex" valid:"email~Email is not valid,required~Email is not valid"`

	SigninID *uint `valid:"-"`
	Signin   Signin `gorm:"references:ID" valid:"-"` 

	Gender_ID *uint  `valid:"required~Please select Gender"`
	Gender    Gender `valid:"-" gorm:"references:id"`

	Province_ID *uint  `valid:"required~Please select Province"`
	Province    Province `valid:"-" gorm:"references:id"`



	Bookings  []Booking   `gorm:"foreignKey:CustomerID"`
	RepairReq []RepairReq `gorm:"foreignKey:CustomerID"`
	Review    []Review    `gorm:"foreignKey:CustomerID"`
	Service   []Service   `gorm:"foreignKey:CustomerID"`
	Payment   []Payment   `gorm:"foreignKey:CustomerID"`
}
