package entity

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	B_name string
	// ส่งไป
	Bookings []Booking `gorm:"foreignKey:BranchID"`
}

type Booking struct {
	gorm.Model
	//รับเข้ามา
	BranchID *uint  `valid:"required~Please select branch"`
	Branch   Branch `valid:"-" gorm:"references:id"`
	//รับเข้ามา
	RoomID *uint `valid:"required~Please select room"`
	Room   Room  `valid:"-" gorm:"references:id"`

	Start time.Time `valid:"required~Please select Start date" govalidator:"func=ValidateStartBeforeStop`
	Stop  time.Time `valid:"required~Please select Stop date" govalidator:"func=ValidateStartBeforeStop`
	//รับเข้ามา
	CustomerID *uint    `valid:"required~Please Signin"`
	Customer   Customer `valid:"-" gorm:"references:id"`

	CheckInOut []CheckInOut `gorm:"foreignKey:BookingID"`
}
